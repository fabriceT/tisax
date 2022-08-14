package evaluation

import (
	"testing"
)

func TestGetQuestion(t *testing.T) {
	LoadYAML("../../tisax.yml")

	catalogs := GetAllCatalogs()

	question := catalogs[0].GetQuestion("7.1.1")
	if question.Isa != "7.1.1" {
		t.Log("Expected 7.1.1 got ", question)
		t.Fail()
	}

	path, err := question.GetQuestionResultPath()
	if err != nil && path != "/tmp/ds.md" {
		t.Log("Expected /tmp/7.1.1.md, got", path)
		t.Fail()
	}

	question = catalogs[1].GetQuestion("7.1.1")
	_, err = question.GetQuestionResultPath()
	if err != nil {
		t.Log("Failed to raise error")
	}

}

func TestGetAllQuestions(t *testing.T) {
	LoadYAML("../../tisax.yml")

	type ResultTest struct {
		name     string
		catalog  int
		expected int
	}

	catalogsTests := []ResultTest{
		{"catalog 1", 0, 41}, // TODO Check: are they 41 questions here ?
		{"chapter 2", 1, 22},
		{"chapter 3", 2, 4},
	}

	catalogs := GetAllCatalogs()

	for _, test := range catalogsTests {
		questions := catalogs[test.catalog].GetAllQuestions()

		t.Run("test",
			func(t *testing.T) {
				t.Log(test.name)

				if len(questions) != test.expected {
					t.Log("Expected ", test.expected, " got ", len(questions))
					t.Fail()
				}
			},
		)

	}
}

func TestLoadEvalution(t *testing.T) {
	LoadYAML("../../tisax.yml")
	catalogs := GetAllCatalogs()

	question := catalogs[0].GetQuestion("7.1.1")
	eval, err := question.LoadResult()
	if err != nil {
		t.Log("Evaluation error", err)
		t.Fail()
	}

	if eval.Note != "1" {
		t.Log("Expected 1, got", eval.Note)
		t.Fail()
	}

	if eval.Note == "" {
		t.Log("incorrect eval.Text", eval.Text)
		t.Fail()
	}
}
