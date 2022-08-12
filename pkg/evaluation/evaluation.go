package evaluation

import (
	"io"
	"log"
	"os"

	"github.com/FabriceT/tisax/pkg/models"
)

var (
	catalogsEntry models.CatalogsEntry
)

func Blabla() int {
	return 1
}

func LoadString(content string) error {
	return nil
}

func LoadFile(filename string) error {

	yaml, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer yaml.Close()

	data, err := io.ReadAll(yaml)
	if err != nil {
		log.Fatal(err)
	}

	if err := catalogsEntry.Parse(data); err != nil {
		log.Fatal(err)
	}

	return nil
}

/*
func GetAllCatalogs() (catalogs []models.CatalogEntry) {
	catalogs = append(catalogs, catalogsEntry.Catalogs...)

	return
}
*/

func GetAllCatalogs() []models.CatalogEntry {
	return catalogsEntry.Catalogs
}

/*
func GetCatalog(isa string) []models.CatalogEntry {
	var catalogs []models.CatalogEntry
	log.Fatal("TODO")

	return catalogs
}

func GetAllChapters() []models.ChaptersEntry {
	var chapters []models.ChaptersEntry

	for _, catalog := range GetAllCatalogs() {
		chapters = append(chapters, catalog.Chapters...)
	}

	return chapters
}

func GetAllChaptersFromCatalog() []models.ChaptersEntry {
	var chapters []models.ChaptersEntry

	log.Fatal("Todo")
	return chapters
}

func GetAllAssessments() []models.AssessmentsEntry {
	var assessments []models.AssessmentsEntry

	for _, chapter := range GetAllChapters() {
		assessments = append(assessments, chapter.Assessments...)
	}

	return assessments
}

func GetAllQuestions() []models.QuestionsEntry {
	var questions []models.QuestionsEntry

	for _, assessments := range GetAllAssessments() {
		questions = append(questions, assessments.Questions...)
	}

	return questions
}
/*

func GetAllQuestionsFromCatalog(isa string) []models.QuestionsEntry {
	var questions []models.QuestionsEntry

	log.Fatal("Todo")
	return questions
}

/*
func GetQuestion(isa string) models.QuestionsEntry {
	for _, catalog := range catalogsEntry.Catalogs {
		for _, chapter := range catalog.Chapters {
			for _, assessment := range chapter.Assessments {
				for _, question := range assessment.Questions {
					if question.Isa == isa {
						return question
					}
				}
			}
		}
	}

	return models.QuestionsEntry{}
}
*/
