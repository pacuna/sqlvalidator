/*
Copyright © 2020 Pablo Acuna <pacuna@pm.me>

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
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pacuna/sqlvalidator/parser/hive"
	"github.com/pacuna/sqlvalidator/parser/mysql"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check if a query contains syntax errors",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch dialect {
		case "hive":
			input := antlr.NewInputStream(args[0])
			lexer := hive.NewHiveLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := hive.NewHiveParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.Statement()
		case "mysql":
			input := antlr.NewInputStream(args[0])
			lexer := mysql.NewMySqlLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := mysql.NewMySqlParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.Root()
		}
	}}

var (
	dialect string
)

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//var dialect string
	validateCmd.PersistentFlags().StringVar(&dialect, "dialect", "", "dialect to validate the query against")
	validateCmd.MarkFlagRequired("dialect")

}
