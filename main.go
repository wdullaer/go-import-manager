// Copyright 2017 Wouter Dullaert
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-import-manager"
	app.Usage = "Reliably manipulate import statements of a go file"
	app.EnableBashCompletion = true
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:      "list",
			Aliases:   []string{"l"},
			Usage:     "list all current imports",
			ArgsUsage: "FILE",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					fmt.Println("list takes exactly one argument: the path of the file to analyze")
					cli.ShowCommandHelpAndExit(c, "list", 1)
					return nil
				}
				file := c.Args()[0]

				imports, err := ListImports(file)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				for _, v := range imports {
					fmt.Println(v)
				}
				return nil
			},
		},
		{
			Name:      "add",
			Aliases:   []string{"a"},
			Usage:     "add an import to the file",
			ArgsUsage: "FILE IMPORT [IMPORT...]",
			Action: func(c *cli.Context) error {
				if c.NArg() < 2 {
					fmt.Println("add takes 2 or more arguments: the file path and 1 or more import statements")
					cli.ShowCommandHelpAndExit(c, "add", 1)
					return nil
				}
				file := c.Args()[0]
				imports := c.Args()[1:]

				str, err := AddImports(file, imports)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				_, err = fmt.Fprint(os.Stdout, str)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

				return nil
			},
		},
		{
			Name:      "delete",
			Aliases:   []string{"d"},
			Usage:     "delete an import from a file",
			ArgsUsage: "FILE IMPORT [IMPORT...]",
			Action: func(c *cli.Context) error {
				if c.NArg() < 2 {
					fmt.Println("delete takes 2 or more arguments: the file path and 1 or more import statements")
					cli.ShowCommandHelpAndExit(c, "delete", 1)
					return nil
				}
				file := c.Args()[0]
				imports := c.Args()[1:]

				str, err := RemoveImports(file, imports)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				_, err = fmt.Fprint(os.Stdout, str)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

				return nil
			},
		},
		{
			Name:      "replace",
			Aliases:   []string{"r"},
			Usage:     "replace an import with another one in a file",
			ArgsUsage: "FILE OLD_IMPORT NEW_IMPORT",
			Action: func(c *cli.Context) error {
				if c.NArg() != 3 {
					fmt.Println("replace takes exactly 3 arguments: the file path, the old import statement and the new import statement")
					cli.ShowCommandHelpAndExit(c, "replace", 1)
					return nil
				}
				file := c.Args()[0]
				oldImport := c.Args()[1]
				newImport := c.Args()[2]

				str, err := ReplaceImport(file, oldImport, newImport)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				_, err = fmt.Fprint(os.Stdout, str)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}
	app.Run(os.Args)
}
