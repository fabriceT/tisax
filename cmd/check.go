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

	"github.com/FabriceT/tisax/pkg/evaluation"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check TISAX evaluation",
	Long:  "Checks if a TISAX evaluation is completed",
	Run: func(cmd *cobra.Command, args []string) {

		evaluation.LoadYAML("tisax.yml")

		catalogs := evaluation.GetAllCatalogs()
		for _, catalog := range catalogs {
			fmt.Println(catalog.Catalog)
			questions := catalog.GetAllQuestions()
			for _, q := range questions {
				// On se fiche de l'erreur
				result, _ := q.LoadResult()
				fmt.Print(" ", q.Isa, ") ", q.Name, " - ", result.Note)

				switch result.Note {
				case "3":
					fmt.Println(" \u2705")
				case "4":
					fmt.Println(" \u2705\u16ED")
				case "5":
					fmt.Println(" \u2705\u16ED\u16ED")
				default:
					fmt.Println(" \u274C")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
