package models

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

type QuestionEntry struct {
	Name      string `yaml:"name"`
	Isa       string `yaml:"isa"`
	Reference string `yaml:"reference"`
	Objective string `yaml:"objective"`
	Must      string `yaml:"must"`
	Should    string `yaml:"should"`
}

type AssessmentsEntry struct {
	Name      string          `yaml:"name"`
	Isa       string          `yaml:"isa"`
	Objective string          `yaml:"objective"`
	Questions []QuestionEntry `yaml:"questions"`
}

type ChaptersEntry struct {
	Chapter     string             `yaml:"chapter"`
	Isa         string             `yaml:"isa"`
	Objective   string             `yaml:"objective"`
	Assessments []AssessmentsEntry `yaml:"assessments"`
}

type CatalogEntry struct {
	Catalog  string          `yaml:"catalog"`
	Chapters []ChaptersEntry `yaml:"chapters"`
}

type CatalogsEntry struct {
	Catalogs []CatalogEntry `yaml:"catalogs"`
}

func (c *CatalogsEntry) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func (c *CatalogEntry) GetQuestion(isa string) (result QuestionEntry) {
	result = QuestionEntry{}

	for _, chapter := range c.Chapters {
		for _, assessment := range chapter.Assessments {
			for _, question := range assessment.Questions {
				if question.Isa == isa {
					result = question
					return
				}
			}
		}
	}

	return
}

func (c *CatalogEntry) GetAllQuestions() (results []QuestionEntry) {
	results = []QuestionEntry{}

	for _, chapter := range c.Chapters {
		results = append(results, chapter.GetAllQuestions()...)
	}

	return
}

func (c *ChaptersEntry) GetAllQuestions() (results []QuestionEntry) {
	results = []QuestionEntry{}
	for _, assessment := range c.Assessments {
		results = append(results, assessment.Questions...)
	}

	return
}

func (q *QuestionEntry) GetQuestionResultPath(evaldir string) (string, error) {
	if q.Isa == "" {
		return "", errors.New("question doesn't contains ISA code")
	}

	return fmt.Sprintf("%s/%s.md", evaldir, strings.ToLower(q.Isa)), nil
}
