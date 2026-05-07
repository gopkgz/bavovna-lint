package analyzers

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"

	"github.com/gopkgz/bavovna-lint/pkg/analyzers/nolinter"
	"github.com/gopkgz/bavovna-lint/pkg/config"
	"github.com/gopkgz/bavovna-lint/pkg/reports"
)

// InspectFunc core logic of the linter, acts on an ast.Node, traversed depth-first
// returns possible lint violations, that will be further inspected for nolint rules.
type InspectFunc func(n ast.Node, importAliases map[string]string, lastPos token.Pos) []reports.Report

// Analyze generates a `run` function for a linter, based on a simple template:
// it does filtering based on configured glob patterns;
// it collects import aliases and sends it as a linter function param
// it collects reports from a linter function and checks them against nolint rules.
func Analyze(inspectFunc InspectFunc) func(pass *analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (any, error) {
		for _, file := range pass.Files {
			if err := processFile(pass, file, inspectFunc); err != nil {
				return nil, err
			}
		}

		return nil, nil
	}
}

// processFile is the per-file pipeline: skip-check, inspect, emit.
func processFile(pass *analysis.Pass, file *ast.File, inspectFunc InspectFunc) error {
	skip, err := config.ShouldSkip(pass.Fset.File(file.Pos()).Name())
	if err != nil {
		return fmt.Errorf("config.ShouldSkip: %w", err)
	}

	if skip {
		return nil
	}

	possibleReports, err := runInspector(file, inspectFunc)
	if err != nil {
		return err
	}

	emitReports(pass, possibleReports)

	return nil
}

// runInspector walks one file, collecting reports and import aliases.
// Returns the collected reports, or an error if inspection fails.
func runInspector(file *ast.File, inspectFunc InspectFunc) ([]*reports.Report, error) {
	possibleReports := []*reports.Report{}
	importAliases := map[string]string{}

	var inspectErr error

	ast.Inspect(file, func(n ast.Node) bool {
		if inspectErr != nil {
			return false
		}

		updateReportPositions(possibleReports, n)

		if err := collectImportAlias(n, importAliases); err != nil {
			inspectErr = err
			return false
		}

		newReports := inspectFunc(n, importAliases, file.End())
		for _, r := range newReports {
			possibleReports = append(possibleReports, &r) //nolint:appendr
		}

		return true
	})

	if inspectErr != nil {
		return nil, inspectErr
	}

	return possibleReports, nil
}

// updateReportPositions advances report end-positions as ast.Inspect walks past them.
func updateReportPositions(possibleReports []*reports.Report, n ast.Node) {
	for _, report := range possibleReports {
		if n != nil && report.Pos < n.Pos() && report.NextTokenPos < n.Pos() {
			report.NextTokenPos = n.Pos()
		}
	}
}

// collectImportAlias adds an import alias to the map if the node is a named ImportSpec.
// Returns an error if the import path cannot be unquoted.
func collectImportAlias(n ast.Node, importAliases map[string]string) error {
	imp, ok := n.(*ast.ImportSpec)
	if !ok || imp.Name == nil {
		return nil
	}

	val, err := strconv.Unquote(imp.Path.Value)
	if err != nil {
		return fmt.Errorf("unquote import path %q: %w", imp.Path.Value, err)
	}

	importAliases[imp.Name.String()] = val

	return nil
}

// emitReports filters reports through nolinter and emits them on the analysis.Pass.
func emitReports(pass *analysis.Pass, possibleReports []*reports.Report) {
	for _, report := range possibleReports {
		if nolinter.IsSupressed(pass, report.Pos, report.NextTokenPos) {
			continue
		}

		pass.Report(analysis.Diagnostic{
			Pos:            report.Pos,
			End:            report.NextTokenPos,
			Category:       report.Category,
			Message:        report.Message,
			URL:            "",
			SuggestedFixes: nil,
			Related:        nil,
		})
	}
}
