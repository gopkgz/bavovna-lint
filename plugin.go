// Package bavovna registers bavovna-lint analyzers with golangci-lint's
// Module Plugin System. Build a custom golangci-lint binary via
// `golangci-lint custom` (config in .custom-gcl.yml) to use them.
package bavovna

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/gopkgz/bavovna-lint/pkg/analyzers/appendr"
	"github.com/gopkgz/bavovna-lint/pkg/analyzers/elser"
	"github.com/gopkgz/bavovna-lint/pkg/analyzers/readall"
)

//nolint:gochecknoinits // register.Plugin must run at package init per golangci-lint plugin contract.
func init() {
	// bavovna bundles all three analyzers under one linter name (back-compat).
	// appendr/elser/readall expose each analyzer as its own linter so that a
	// per-analyzer nolint directive is recognized by golangci-lint and findings
	// are attributed to that name. Enable EITHER the bundle OR the individual
	// linters — never both, or the same check runs under two linter names.
	register.Plugin("bavovna", newPlugin(appendr.Analyzer, elser.Analyzer, readall.Analyzer))
	register.Plugin("appendr", newPlugin(appendr.Analyzer))
	register.Plugin("elser", newPlugin(elser.Analyzer))
	register.Plugin("readall", newPlugin(readall.Analyzer))
}

// newPlugin returns a constructor exposing the given analyzers as one linter.
// Settings are unused — analyzers ship their own flags.
func newPlugin(analyzers ...*analysis.Analyzer) register.NewPlugin {
	return func(_ any) (register.LinterPlugin, error) {
		return &bavovnaPlugin{analyzers: analyzers}, nil
	}
}

type bavovnaPlugin struct {
	analyzers []*analysis.Analyzer
}

func (b *bavovnaPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return b.analyzers, nil
}

func (b *bavovnaPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
