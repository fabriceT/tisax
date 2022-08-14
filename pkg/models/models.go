package models

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type QuestionEntry struct {
	Name       string `yaml:"name"`
	Isa        string `yaml:"isa"`
	Reference  string `yaml:"reference"`
	Objectives string `yaml:"objectives"`
	Must       string `yaml:"must"`
	Should     string `yaml:"should"`
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
		for _, assessment := range chapter.Assessments {
			results = append(results, assessment.Questions...)
		}
	}

	return
}

func (q *QuestionEntry) GetQuestionResultPath() (string, error) {
	if q.Isa == "" {
		return "", errors.New("question doesn't contains ISA code")
	}

	return fmt.Sprintf("./evaluation/%s.md", strings.ToLower(q.Isa)), nil
}

type EvaluationResult struct {
	Note string
	Text string
}

func (q *QuestionEntry) LoadResult() (EvaluationResult, error) {
	eval := EvaluationResult{
		Note: "0",
		Text: "",
	}

	re, err := regexp.Compile(`NOTE=(\d)`)
	if err != nil {
		panic("Woops")
	}

	path, err := q.GetQuestionResultPath()
	if err != nil {
		return eval, err
	}

	file, err := os.Open(path)
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
