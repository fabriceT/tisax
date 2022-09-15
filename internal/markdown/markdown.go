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
	fmt.Fprintf(&md, "### %s) %s {#%s}\n", chapter.Isa, chapter.Chapter, chapter.Isa)
}

func AddQuestion(question models.QuestionEntry, maturityLevel int64, text string) {

	// On empêche l'affichage de valeur négative (pas d'évaluation)
	if maturityLevel < 0 {
		maturityLevel = 0
	}

	fmt.Fprintf(&md, "#### %s) %s {#%s}\n\n", question.Isa, question.Name, question.Isa)

	if question.Must != "" {
		fmt.Fprintf(&md, "{.must .evaltext%d}\n%s\n", maturityLevel, question.Must)
	}

	if question.Should != "" {
		fmt.Fprintf(&md, "{.should}\n%s\n", question.Should)
	}

	if question.Objective != "" {
		fmt.Fprintf(&md, "{.objective}\n%s\n\n", question.Objective)
	}

	fmt.Fprint(&md, "\n")

	if text != "" {
		in := strings.ReplaceAll(text, "%NAME%", question.Name)
		in = strings.ReplaceAll(in, "%ISA%", question.Isa)
		in = strings.ReplaceAll(in, "%REFERENCE%", question.Reference)
		in = strings.ReplaceAll(in, "%OBJECTIVE%", question.Reference)
		in = strings.ReplaceAll(in, "%MUST%", question.Must)
		in = strings.ReplaceAll(in, "%SHOULD%", question.Should)
		in = strings.ReplaceAll(in, "%MATURITY_LEVEL%", fmt.Sprintf("%d", maturityLevel))
		in = strings.ReplaceAll(in, "%MATURITY_ICON%", internal.GetMaturityIcon(maturityLevel))

		fmt.Fprintf(&md, "%s\n", in)
	} else {
		md.WriteString("Non évalué\n")
	}
}

func IncludeMDFile(filename string) {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Printf("Warning: %s, file %s not included", err.Error(), filename)
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

	f2, _ := os.Create("cr.md")
	defer f2.Close()

	f2.Write([]byte(md.String()))

}
