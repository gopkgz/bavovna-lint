package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/gopkgz/bavovna-lint/pkg/analyzers/appendr"
	"github.com/gopkgz/bavovna-lint/pkg/analyzers/elser"
	"github.com/gopkgz/bavovna-lint/pkg/analyzers/readall"
)

func main() {
	unitchecker.Main(
		appendr.Analyzer,
		elser.Analyzer,
		readall.Analyzer,
	)
}
