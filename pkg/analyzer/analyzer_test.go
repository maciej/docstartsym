package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	// Get the path to testdata directory
	testdata := analysistest.TestData()

	// Run the analyzer on the test files in testdata/src/a/
	// The analyzer will check if documentation comments start with the symbol name
	// and report any issues found
	analysistest.Run(t, testdata, Analyzer, "a")
}
