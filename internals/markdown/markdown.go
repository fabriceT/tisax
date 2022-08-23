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
	extensions := parser.CommonExtensions
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
		md += fmt.Sprintf("* <span class=eval%s><u>Évaluation</u> : %s</span> %s\n\n", note, note, internals.GetNoteIcon(note))
		md += fmt.Sprintf("<p> <span class=remarque>Remarque</span> :</p><div>%s</div>\n", text)

	} else {
		md += "\nNon évalué\n"
	}

}

func AddLine() {
	md += "***\n"
}

func Save() {
	f, err := os.Create("evaluation.html")

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
