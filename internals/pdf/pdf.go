package pdf

import (
	"github.com/FabriceT/tisax/internals/models"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

var m pdf.Maroto

func init() {
	m = pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
}

func AddCatalog(catalog models.CatalogEntry) {
	m.Row(15, func() {
		m.Col(12, func() {
			m.Text("Catalog: "+catalog.Catalog, props.Text{
				Size:  16,
				Style: consts.Bold,
			})
		})
	})
}

func AddChapter(chapter models.ChaptersEntry) {
	m.Row(14, func() {
		m.Col(12, func() {
			m.Text(chapter.Isa+") "+chapter.Chapter, props.Text{
				Size:  14,
				Style: consts.Bold,
			})
		})
	})

}

func AddQuestion(question models.QuestionEntry, note string, text string) {

	m.Row(8, func() {
		m.Col(1, func() {
			m.Text(question.Isa, props.Text{
				Style: consts.Bold,
			})
		})
		m.Col(11, func() {
			m.Text(question.Name, props.Text{
				Style: consts.Bold,
			})
		})
	})

	tableHeadings := []string{"Item", "Description"}
	contents := [][]string{}

	if question.Objective != "" {
		contents = append(contents, []string{"Objective", question.Objective})
	}
	contents = append(contents, []string{"Must", question.Must})

	if note != "" {
		contents = append(contents, []string{"Note", note})
	}
	contents = append(contents, []string{"Evaluation", text})

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 10},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 10},
		},
		Align:              consts.Left,
		HeaderContentSpace: 1,
		Line:               false,
	})

	m.Row(8, func() {})
}

func Save(filename string) error {
	return m.OutputFileAndClose(filename)
}
