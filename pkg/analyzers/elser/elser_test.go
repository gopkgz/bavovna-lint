package elser

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestElser(t *testing.T) {
	t.Parallel()

	testdata, err := filepath.Abs("../../../testdata/elser")
	if err != nil {
		log.Fatal(err)
	}

	res := analysistest.Run(t, testdata, Analyzer, "")
	require.NotNil(t, res)
	require.NotEmpty(t, res)
	require.Len(t, res[0].Diagnostics, 1)

	diagRes := res[0].Diagnostics[0]
	require.Equal(t, analyzerName, diagRes.Category)
	require.Equal(t, analyzerMsg, diagRes.Message)
}
