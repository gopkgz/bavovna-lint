package appendr

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/gopkgz/bavovna-lint/pkg/analyzers"
	"github.com/gopkgz/bavovna-lint/pkg/analyzers/nolinter"
	"github.com/gopkgz/bavovna-lint/pkg/reports"
)

const analyzerName = "appendr"
const analyzerMsg = "append is not efficient on the heap and is not prone to race conditions"

// Analyzer append is not efficient on the heap and is not prone to race conditions.
//
//nolint:gochecknoglobals,exhaustruct // exported Analyzer per analysis package convention; analysis.Analyzer.Flags zero value (flag.FlagSet) is the documented default.
var Analyzer = &analysis.Analyzer{
	Name:     analyzerName,
	Doc:      "finds append statements in the code",
	Run:      analyzers.Analyze(run),
	Requires: []*analysis.Analyzer{nolinter.Analyzer},
}

func run(n ast.Node, importAliases map[string]string, lastPos token.Pos) []reports.Report {
	res := []reports.Report{}

	if call, ok := n.(*ast.CallExpr); ok {
		if ident, ok := call.Fun.(*ast.Ident); ok {
			if ident.Name == "append" {
				res = append(res, reports.Report{ //nolint:appendr
					Pos:          ident.Pos(),
					NextTokenPos: lastPos,
					Category:     analyzerName,
					Message:      analyzerMsg,
				})
			}
		}
	}

	return res
}
