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
