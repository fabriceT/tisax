/*
TISAX Evaluation : mardown to HTML
Copyright (C) 2022 Fabrice THIROUX

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/
*/

package models

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
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

type EvaluationResult struct {
	MaturityLevel int64
	Text          string
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

func (q *QuestionEntry) GetResult(basepath string) (EvaluationResult, error) {
	eval := EvaluationResult{
		MaturityLevel: -1,
		Text:          "",
	}

	if q.Isa == "" {
		return eval, errors.New("question doesn't contains ISA code")
	}

	filename := path.Join(basepath, strings.ToLower(q.Isa)+".md")

	re, err := regexp.Compile(`NOTE=(\d)`)
	if err != nil {
		panic("Woops")
	}

	file, err := os.Open(filename)
	if err != nil {
		return eval, err
	}
	defer file.Close()

	eval.Text = ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if re.Match([]byte(line)) {
			results := re.FindStringSubmatch(line)
			// TODO Add check
			eval.MaturityLevel, _ = strconv.ParseInt(results[1], 10, 4)
		} else {
			eval.Text += line + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return eval, nil

}
