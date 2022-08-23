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
	"github.com/FabriceT/tisax/internals/evaluation"
	"github.com/FabriceT/tisax/internals/pdf"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export tisax report to a PDF File",
	Long:  "Write a evaluation report to a PDF file",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO Export TISAX evaluation into a document (markdown ?)

		evaldir, _ := cmd.Flags().GetString("evaldir")
		yamlFile, _ := cmd.Flags().GetString("file")

		evaluation.LoadYAML(yamlFile)

		catalogs := evaluation.GetAllCatalogs()
		for _, catalog := range catalogs {
			pdf.AddCatalog(catalog)
			for _, chapter := range catalog.Chapters {
				pdf.AddChapter(chapter)
				for _, assessment := range chapter.Assessments {
					for _, question := range assessment.Questions {
						path, _ := question.GetQuestionResultPath(evaldir)
						result, _ := evaluation.LoadEvaluationResult(path)
						pdf.AddQuestion(question, result.Note, result.Text)
					}
					pdf.AddLine()
				}
				/*
					for _, question := range chapter.GetAllQuestions() {
						path, _ := question.GetQuestionResultPath(evaldir)
						result, _ := evaluation.LoadEvaluationResult(path)
						pdf.AddQuestion(question, result.Note, result.Text)
					}
				*/

			}
		}

		pdf.Save("evaluation.pdf")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
