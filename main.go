package main

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var forbiddenImports = []string{
	// Add your forbidden imports here
	"path",
}

var Analyzer = &analysis.Analyzer{
	Name: "badimports",
	Doc:  "Checks for forbidden package imports",
	Run:  run,
}

func main() {
	singlechecker.Main(Analyzer)
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			importDecl, ok := n.(*ast.ImportSpec)
			if !ok {
				return true
			}

			importPath := importDecl.Path.Value
			for _, forbiddenImport := range forbiddenImports {
				if "\""+forbiddenImport+"\"" == importPath {
					pass.Reportf(importDecl.Pos(), "Forbidden import: %s", forbiddenImport)
				}
			}

			return true
		})
	}

	return nil, nil
}
