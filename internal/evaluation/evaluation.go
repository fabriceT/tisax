package evaluation

import (
	"io"
	"log"
	"os"

	"github.com/FabriceT/tisax/internal/models"
)

var (
	catalogsEntry models.CatalogsEntry
)

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
