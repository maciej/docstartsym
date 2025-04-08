package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "docstartsym",
	Doc:  "check that doc comments start with the symbol name",
	URL:  "https://github.com/maciej/docstartsym",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				if d.Doc == nil {
					continue
				}

				// Get the first line of the documentation
				docText := d.Doc.Text()
				firstLine := strings.Split(docText, "\n")[0]

				// Check if the first line starts with the function name
				if !strings.HasPrefix(strings.TrimSpace(firstLine), d.Name.Name) {
					pass.Reportf(d.Doc.Pos(),
						"doc comment should start with the function name %q",
						d.Name.Name)
				}

			case *ast.GenDecl:
				if d.Doc == nil {
					continue
				}

				// Get the first line of the documentation
				docText := d.Doc.Text()
				firstLine := strings.Split(docText, "\n")[0]

				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec:
						// Handle type declarations
						if !strings.HasPrefix(strings.TrimSpace(firstLine), s.Name.Name) {
							pass.Reportf(d.Doc.Pos(),
								"doc comment should start with the type name %q",
								s.Name.Name)
						}

					case *ast.ValueSpec:
						// Handle var and const declarations
						for _, name := range s.Names {
							if !strings.HasPrefix(strings.TrimSpace(firstLine), name.Name) {
								pass.Reportf(d.Doc.Pos(),
									"doc comment should start with the %s name %q",
									getDeclType(d.Tok), name.Name)
							}
						}
					}
				}
			}
		}
	}

	return nil, nil
}

func getDeclType(tok token.Token) string {
	switch tok {
	case token.VAR:
		return "variable"
	case token.CONST:
		return "constant"
	default:
		return "symbol"
	}
}
