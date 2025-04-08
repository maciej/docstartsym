package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "docstartsym",
	Doc:  "check that documentation starts with the symbol name",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				if funcDecl.Doc == nil {
					continue
				}

				// Get the first line of the documentation
				docText := funcDecl.Doc.Text()
				firstLine := strings.Split(docText, "\n")[0]

				// Check if the first line starts with the function name
				if !strings.HasPrefix(strings.TrimSpace(firstLine), funcDecl.Name.Name) {
					pass.Reportf(funcDecl.Doc.Pos(),
						"documentation should start with the function name %q",
						funcDecl.Name.Name)
				}
			}
		}
	}

	return nil, nil
}
