package main

import (
	"github.com/Alandres998/url-shortner/cmd/staticlint/myanalyzers"
	"github.com/kisielk/errcheck/errcheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

// initAnalyzers инициализация анализаторов
func initAnalyzers() []*analysis.Analyzer {
	var analyzers []*analysis.Analyzer

	//Анализаторы из коробочки
	analyzers = append(analyzers, nilfunc.Analyzer, shadow.Analyzer, structtag.Analyzer)

	for _, v := range staticcheck.Analyzers {
		if v.Analyzer.Name == "SA" || v.Analyzer.Name == "S1000" {
			analyzers = append(analyzers, v.Analyzer)
		}
	}

	// Добавляем публичный анализатор errcheck
	analyzers = append(analyzers, errcheck.Analyzer)

	// Кастомный анализатор
	analyzers = append(analyzers, myanalyzers.ProhibitOsExitInMainAnalyzer)

	return analyzers
}

// main запускает multichecker с инициализированными анализаторами.
func main() {
	analyzers := initAnalyzers()
	multichecker.Main(analyzers...)
}
