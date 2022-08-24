package markdown

import (
	"fmt"
	"log"
	"os"

	"github.com/FabriceT/tisax/internals"
	"github.com/FabriceT/tisax/internals/models"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	md string = ""
	p  *parser.Parser
)

func init() {
	extensions := parser.CommonExtensions | parser.Footnotes
	p = parser.NewWithExtensions(extensions)
}

func AddCatalog(catalog models.CatalogEntry) {
	md += fmt.Sprintf("## %s\n", catalog.Catalog)
}

func AddChapter(chapter models.ChaptersEntry) {
	md += fmt.Sprintf("### %s) %s\n", chapter.Isa, chapter.Chapter)
}

func AddQuestion(question models.QuestionEntry, note string, text string) {
	md += fmt.Sprintf("#### %s) %s\n", question.Isa, question.Name)

	if question.Objective != "" {
		//md += fmt.Sprintf("<div class=objective>%s</div>\n\n", question.Objective)
		md += fmt.Sprintf("* <span>Objectif</span> : %s\n", question.Objective)
	}

	if question.Reference != "" {
		md += fmt.Sprintf("* <span>Référence</span> : %s\n", question.Reference)
	}

	if text != "" {
		md += fmt.Sprintf("* <span>Évaluation</span> : %s %s\n\n", note, internals.GetNoteIcon(note))
		md += fmt.Sprintf("<div class='evaltext%s'>%s</div>\n", note, text)

	} else {
		md += "\nNon évalué\n"
	}

}

func AddLine() {
	md += "***\n"
}

func Save(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage | html.TOC
	opts := html.RendererOptions{
		Flags: htmlFlags,
		Title: "Evaluation TISAX",
		CSS:   "style.css",
	}
	renderer := html.NewRenderer(opts)
	html := markdown.ToHTML([]byte(md), p, renderer)

	_, err2 := f.Write(html)
	if err2 != nil {
		log.Fatal(err2)
	}
}
