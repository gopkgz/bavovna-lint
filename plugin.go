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
	register.Plugin("bavovna", New)
}

// New constructs the plugin. Settings are unused — analyzers ship their own flags.
//
//nolint:ireturn // register.LinterPlugin return type fixed by golangci-lint plugin contract.
func New(_ any) (register.LinterPlugin, error) {
	return &bavovnaPlugin{}, nil
}

type bavovnaPlugin struct{}

func (b *bavovnaPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		appendr.Analyzer,
		elser.Analyzer,
		readall.Analyzer,
	}, nil
}

func (b *bavovnaPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
