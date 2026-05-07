package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/gopkgz/bavovna-lint/pkg/analyzers/readall"
)

func main() {
	unitchecker.Main(
		readall.Analyzer,
	)
}
