package evaluation

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/FabriceT/tisax/pkg/models"
)

var (
	catalogsEntry models.CatalogsEntry
)

type EvaluationResult struct {
	Note string
	Text string
}

func LoadString(content string) error {
	return nil
}

func LoadYAML(filename string) error {

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

func LoadEvaluationResult(filename string) (EvaluationResult, error) {
	eval := EvaluationResult{
		Note: "0",
		Text: "",
	}

	re, err := regexp.Compile(`NOTE=(\d)`)
	if err != nil {
		panic("Woops")
	}

	file, err := os.Open(filename)
	if err != nil {
		return eval, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if re.Match([]byte(line)) {
			results := re.FindStringSubmatch(line)
			// TODO Add check
			eval.Note = string(results[1])
		} else {
			eval.Text += line + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return eval, nil
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
