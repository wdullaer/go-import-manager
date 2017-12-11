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
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
)

func ListImports(fname string) ([]string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fname, nil, parser.ImportsOnly)

	if err != nil {
		return nil, err
	}

	result := make([]string, len(node.Imports))
	for i, s := range node.Imports {
		result[i] = s.Path.Value
	}

	return result, nil
}

func AddImports(fname string, imports []string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)

	if err != nil {
		return "", err
	}

	ast.Inspect(node, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if ok && decl.Tok == token.IMPORT {
			// Add all the imports
			for _, v := range imports {
				decl.Specs = append(decl.Specs, &ast.ImportSpec{Path: &ast.BasicLit{Value: ensureQuotes(v)}})
			}
			return false
		}
		return true
	})

	ast.SortImports(fset, node)

	return astToString(fset, node)
}

func RemoveImports(fname string, imports []string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)

	if err != nil {
		return "", err
	}

	oldImports := make([]string, len(node.Imports))
	for i, v := range node.Imports {
		oldImports[i] = v.Path.Value
	}

	for i, imp := range imports {
		imp = ensureQuotes(imp)
		imports[i] = imp
		if !includes(oldImports, imp) {
			fmt.Println("WARN: Could not find import", imp)
		}
	}

	ast.Inspect(node, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if ok && decl.Tok == token.IMPORT {
			var newSpecs []ast.Spec
			for _, v := range decl.Specs {
				iSpec, _ := v.(*ast.ImportSpec)
				if !includes(imports, iSpec.Path.Value) {
					newSpecs = append(newSpecs, iSpec)
				}
			}
			decl.Specs = newSpecs
			return false
		}
		return true
	})

	ast.SortImports(fset, node)

	return astToString(fset, node)
}

func ReplaceImport(fname string, oldImport string, newImport string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}
	oldImport = ensureQuotes(oldImport)
	newImport = ensureQuotes(newImport)

	imports := make([]string, len(node.Imports))
	for i, v := range node.Imports {
		imports[i] = v.Path.Value
	}

	if !includes(imports, oldImport) {
		return "", errors.New(fmt.Sprint("WARN Could not find import statement ", oldImport))
	}

	ast.Inspect(node, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if ok && decl.Tok == token.IMPORT {
			for i, v := range decl.Specs {
				iSpec, _ := v.(*ast.ImportSpec)
				if oldImport == iSpec.Path.Value {
					decl.Specs[i] = &ast.ImportSpec{Path: &ast.BasicLit{Value: newImport}}
				}
			}
			return false
		}
		return true
	})

	ast.SortImports(fset, node)

	return astToString(fset, node)
}

func includes(arr []string, item string) bool {
	for _, v := range arr {
		if item == v {
			return true
		}
	}
	return false
}

func astToString(fset *token.FileSet, node interface{}) (string, error) {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, node)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func ensureQuotes(str string) string {
	if strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"") {
		return str
	}
	if (strings.HasPrefix(str, "_ ")) {
		return str
	}
	return strconv.Quote(str)
}
