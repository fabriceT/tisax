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

	"github.com/FabriceT/tisax/internals"
	"github.com/FabriceT/tisax/internals/evaluation"
	"github.com/FabriceT/tisax/internals/markdown"
	"github.com/spf13/cobra"
)

const resultCatalogHeader = `

### %s

| Question | Maturité |
|----------|----------|
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

		| Question    | Maturité |
		|-------------|----------|
		| Automation? |    0     |
		*/
		resultsTextMD := "## Synthèse\n"

		catalogs := evaluation.GetAllCatalogs()
		for _, catalog := range catalogs {

			markdown.AddCatalog(catalog)
			// Results table starts
			resultsTextMD += fmt.Sprintf(resultCatalogHeader, catalog.Catalog)

			for _, chapter := range catalog.Chapters {
				markdown.AddChapter(chapter)
				for _, assessment := range chapter.Assessments {
					for _, question := range assessment.Questions {
						path, _ := question.GetQuestionResultPath(evaldir)
						result, _ := evaluation.LoadEvaluationResult(path)
						markdown.AddQuestion(question, result.MaturityLevel, result.Text)

						// Add item in results table
						resultsTextMD += fmt.Sprintf("| %s - %s | %d %s |\n",
							question.Isa,
							question.Name,
							result.MaturityLevel,
							internals.GetMaturityIcon((result.MaturityLevel)))
					}
					markdown.AddLine()
				}
			}
		}

		// On inclus un fichier conclusion.md si présent.
		markdown.IncludeMDFile(evaldir + "/conclusion.md")

		if synthesis {
			// On inclus la Synthese
			markdown.IncludeMDContent(resultsTextMD)
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
