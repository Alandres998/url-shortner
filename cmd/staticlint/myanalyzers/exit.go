package myanalyzers

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// MainExitAnalyzer — анализатор, который запрещает использование os.Exit в функции main пакета main.
var ProhibitOsExitInMainAnalyzer = &analysis.Analyzer{
	Name: "mainexit",
	Doc:  "запрещает использование os.Exit в функции main пакета main",
	Run:  run,
}

// run функция запуска проверки
func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {
		// Ограничения на мейн
		if pass.Pkg.Name() != "main" {
			continue
		}

		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name.Name != "main" {
				continue
			}

			ast.Inspect(fn.Body, func(n ast.Node) bool {
				if call, ok := n.(*ast.CallExpr); ok {
					if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
						if pkg, ok := sel.X.(*ast.Ident); ok && pkg.Name == "os" && sel.Sel.Name == "Exit" {
							pass.Reportf(call.Pos(), "использование os.Exit в функции main запрещено")
						}
					}
				}
				return true
			})
		}
	}

	return nil, nil
}
