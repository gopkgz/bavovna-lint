package bavovna

import (
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/stretchr/testify/require"
)

func TestRegisteredPlugins(t *testing.T) {
	t.Parallel()

	cases := map[string][]string{
		"bavovna": {"appendr", "elser", "readall"},
		"appendr": {"appendr"},
		"elser":   {"elser"},
		"readall": {"readall"},
	}
	for name, want := range cases {
		ctor, err := register.GetPlugin(name)
		require.NoError(t, err)

		plug, err := ctor(nil)
		require.NoError(t, err)
		require.Equal(t, register.LoadModeSyntax, plug.GetLoadMode())

		got, err := plug.BuildAnalyzers()
		require.NoError(t, err)

		names := make([]string, len(got))
		for i, a := range got {
			names[i] = a.Name
		}

		require.ElementsMatch(t, want, names)
	}
}
