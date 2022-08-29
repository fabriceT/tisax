/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/FabriceT/tisax/internal"
	"github.com/FabriceT/tisax/internal/evaluation"
	"github.com/FabriceT/tisax/internal/markdown"
	"github.com/spf13/cobra"
)

const resultCatalogHeader = `

### %s

 ISA | Question | Maturité
-----|----------|---------
`

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export tisax report to a PDF File",
	Long:  "Write a evaluation report to a PDF file",
	Run: func(cmd *cobra.Command, args []string) {
		evaldir, _ := cmd.Flags().GetString("evaldir")
		yamlFile, _ := cmd.Flags().GetString("file")

		evaluation.LoadYAML(yamlFile)

		/* Markdown content for results table

		## Résultats

		### Catalog

		ISA | Question | Maturité
		----|----------|---------
		X.X | What?    |   0  :D
		*/

		var synthesisBuilder strings.Builder
		synthesisBuilder.WriteString("## Synthèse\n")

		// We add head.md if it exists.
		markdown.IncludeMDFile(evaldir + "/head.md")

		catalogs := evaluation.GetAllCatalogs()
		for _, catalog := range catalogs {

			markdown.AddCatalog(catalog)
			// Results table starts
			fmt.Fprintf(&synthesisBuilder, resultCatalogHeader, catalog.Catalog)

			for _, chapter := range catalog.Chapters {
				markdown.AddChapter(chapter)
				for _, assessment := range chapter.Assessments {
					for _, question := range assessment.Questions {
						path, _ := question.GetQuestionResultPath(path.Join(evaldir, catalog.Catalog))
						result, _ := evaluation.LoadEvaluationResult(path)
						markdown.AddQuestion(question, result.MaturityLevel, result.Text)

						// Add item in results table
						if result.Text == "" {
							// Not evaluated
							fmt.Fprintf(&synthesisBuilder, "%s | %s | -\n",
								question.Isa,
								question.Name)
						} else {
							// | ISA) Question | Note Icon |
							fmt.Fprintf(&synthesisBuilder, "%s | %s | %d %s\n",
								question.Isa,
								question.Name,
								result.MaturityLevel,
								internal.GetMaturityIcon(result.MaturityLevel))
						}
					}
					markdown.AddLine()
				}
			}
		}

		markdown.IncludeMDFile(evaldir + "/end.md")

		if synthesis {
			// On inclus la Synthese
			markdown.IncludeMDContent(synthesisBuilder.String())
		}

		markdown.Save(outputfile)
	},
}
var outputfile string
var synthesis bool

func init() {
	rootCmd.PersistentFlags().StringVar(&outputfile, "output", "evaluation.html", "where to store HTML result")
	rootCmd.PersistentFlags().BoolVar(&synthesis, "synthesis", true, "Add synthesis to document")
	rootCmd.AddCommand(exportCmd)
}
