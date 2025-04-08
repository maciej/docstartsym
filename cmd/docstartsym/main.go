package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/maciej/docstartsym/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
