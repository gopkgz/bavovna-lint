package elser

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/gopkgz/bavovna-lint/pkg/analyzers"
	"github.com/gopkgz/bavovna-lint/pkg/reports"
)

const analyzerName = "elser"
const analyzerMsg = "else is unnecessary, prefer early termination"

// Analyzer else is unnecessary, prefer early termination.
//
//nolint:gochecknoglobals,exhaustruct // exported Analyzer per analysis package convention; analysis.Analyzer.Flags zero value (flag.FlagSet) is the documented default.
var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  "finds else statements in the code",
	Run:  analyzers.Analyze(run),
}

func run(n ast.Node, importAliases map[string]string, lastPos token.Pos) []reports.Report {
	res := []reports.Report{}

	if b, ok := n.(*ast.IfStmt); ok {
		if b.Else != nil {
			res = append(res, reports.Report{ //nolint:appendr
				Pos:          b.Else.Pos(),
				NextTokenPos: lastPos,
				Category:     analyzerName,
				Message:      analyzerMsg,
			})
		}
	}

	return res
}
