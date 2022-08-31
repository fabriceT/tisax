package markdown

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/FabriceT/tisax/internal"
	"github.com/FabriceT/tisax/internal/models"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	md strings.Builder
	p  *parser.Parser
)

func init() {
	extensions := parser.CommonExtensions | parser.Footnotes | parser.Attributes | parser.Mmark
	p = parser.NewWithExtensions(extensions)
}

func AddCatalog(catalog models.CatalogEntry) {
	fmt.Fprintf(&md, "## %s\n", catalog.Catalog)
}

func AddChapter(chapter models.ChaptersEntry) {
	fmt.Fprintf(&md, "### %s) %s\n", chapter.Isa, chapter.Chapter)
}

func AddQuestion(question models.QuestionEntry, maturityLevel int64, text string) {

	// On empêche l'affichage de valeur négative (pas d'évaluation)
	if maturityLevel < 0 {
		maturityLevel = 0
	}

	fmt.Fprintf(&md, "#### %s) %s\n", question.Isa, question.Name)

	if question.Objective != "" {
		fmt.Fprintf(&md, "* <span>Objectif</span> : %s\n", question.Objective)
	}

	if question.Reference != "" {
		fmt.Fprintf(&md, "* <span>Référence</span> : %s\n", question.Reference)
	}

	if text != "" {
		fmt.Fprintf(&md, "* <span>Niveau maturité</span> : %d %s\n\n",
			maturityLevel,
			internal.GetMaturityIcon(maturityLevel))
		fmt.Fprintf(&md, "<div class='eval evaltext%d'>\n%s</div>\n",
			maturityLevel,
			text)
	} else {
		md.WriteString("\nNon évalué\n")
	}
}

func IncludeMDFile(filename string) {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Printf("Error: %s, file %s not included", err.Error(), filename)
	}

	// On ajoute un saut de ligne pour ne pas interférer avec la suite du markdown
	md.Write(content)
	md.WriteByte('\n')
}

func IncludeMDContent(content string) {
	md.WriteString(content)
}

func AddLine() {
	md.WriteString("***\n")
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
	html := markdown.ToHTML([]byte(md.String()), p, renderer)

	_, err2 := f.Write(html)
	if err2 != nil {
		log.Fatal(err2)
	}
}
