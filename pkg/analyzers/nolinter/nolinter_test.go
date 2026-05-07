package nolinter

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNolints(t *testing.T) {
	t.Parallel()

	testdata, err := filepath.Abs("../../../testdata/nolinter")
	if err != nil {
		log.Fatal(err)
	}

	res := analysistest.Run(t, testdata, Analyzer, "")
	require.NotNil(t, res)
	require.NotEmpty(t, res)

	facts, ok := res[0].Result.([]NolintComment)
	require.True(t, ok)
	require.Len(t, facts, 5)
}
