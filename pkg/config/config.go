package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gobwas/glob"
)

const ignorePathsEnv = "IGNORE"

//nolint:gochecknoglobals // env-var IGNORE cache; refactor candidate (Skipper struct) deferred to keep linter-rollout change minimal.
var ignorePaths []string

//nolint:gochecknoglobals // sync.Once guarding ignorePaths init; refactor candidate (Skipper struct) deferred to keep linter-rollout change minimal.
var once sync.Once

// ShouldSkip checks `ignorePathsEnv` environment variable for glob patterns and
// matches those against the given `filename`.
func ShouldSkip(filename string) (bool, error) {
	once.Do(func() {
		ignorePaths = strings.Split(os.Getenv(ignorePathsEnv), ",")
	})

	for _, p := range ignorePaths {
		g, err := glob.Compile(p, '/')
		if err != nil {
			return false, fmt.Errorf("compile ignore-path glob %q: %w", p, err)
		}

		if g.Match(filename) {
			return true, nil
		}
	}

	return false, nil
}
