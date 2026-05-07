package readall

import (
	"fmt"
	"go/ast"
	"go/token"
	"path"

	"golang.org/x/tools/go/analysis"

	"github.com/gopkgz/bavovna-lint/pkg/analyzers"
	"github.com/gopkgz/bavovna-lint/pkg/analyzers/nolinter"
	"github.com/gopkgz/bavovna-lint/pkg/reports"
)

const analyzerName = "readall"
const analyzerMsg = "ioutil.ReadAll is expensive and should be avoided"

// Analyzer ioutil.ReadAll is expensive. This linter will nudge you about ioutil.ReadAll presence in your code.
//
//nolint:gochecknoglobals,exhaustruct // exported Analyzer per analyzer-tool convention (singlechecker / vettool); analysis.Analyzer.Flags zero value (flag.FlagSet) is the documented default.
var Analyzer = &analysis.Analyzer{
	Name:     analyzerName,
	Doc:      "finds ioutil.ReadAll usages",
	Run:      analyzers.Analyze(run),
	Requires: []*analysis.Analyzer{nolinter.Analyzer},
}

func run(n ast.Node, importAliases map[string]string, lastPos token.Pos) []reports.Report {
	pkgID, funcName, ok := selectorIdent(n)
	if !ok {
		return []reports.Report{}
	}

	pkgName := pkgID.Name
	if alias, ok := importAliases[pkgName]; ok {
		pkgName = path.Base(alias)
	}

	if fmt.Sprintf("%s.%s", pkgName, funcName) != "ioutil.ReadAll" {
		return []reports.Report{}
	}

	return []reports.Report{{
		Pos:          pkgID.Pos(),
		NextTokenPos: lastPos,
		Category:     analyzerName,
		Message:      analyzerMsg,
	}}
}

// selectorIdent extracts the package identifier and function name from a
// `pkg.Func(...)` call expression. Returns ok=false for any other node shape.
func selectorIdent(n ast.Node) (*ast.Ident, string, bool) {
	call, ok := n.(*ast.CallExpr)
	if !ok {
		return nil, "", false
	}

	fun, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, "", false
	}

	pkgID, ok := fun.X.(*ast.Ident)
	if !ok {
		return nil, "", false
	}

	return pkgID, fun.Sel.Name, true
}
