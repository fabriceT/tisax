package evaluation

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/FabriceT/tisax/internals/models"
)

var (
	catalogsEntry models.CatalogsEntry
)

type EvaluationResult struct {
	MaturityLevel int64
	Text          string
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

func GetAllCatalogs() []models.CatalogEntry {
	return catalogsEntry.Catalogs
}

func LoadEvaluationResult(filename string) (EvaluationResult, error) {
	eval := EvaluationResult{
		MaturityLevel: 0,
		Text:          "",
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
