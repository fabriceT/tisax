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

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	yamlFile            string
	evaluationDirectory string
)

var rootCmd = &cobra.Command{
	Use:   "tisax",
	Short: "Manage TISAX evalutation",
	Long:  "Do something with TISAX evalution.",
	//Run: func(cmd *cobra.Command, args []string) { fmt.Println("Hello from root cmd") },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&evaluationDirectory, "evaldir", "./evaluation", "evaluation directory (default ./evaluation")
	rootCmd.PersistentFlags().StringVar(&yamlFile, "file", "tisax.yml", "evaluation file (default tisax.yml")
}
