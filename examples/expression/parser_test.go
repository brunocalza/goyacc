package expression

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	tests := []struct {
		expression string
		expected   float32
	}{
		{
			"2",
			2,
		},
		{
			"2 + (3 + 4)",
			9,
		},
		{
			"((((3) + 3)))",
			6,
		},
		{
			"(5 - 3) - 2",
			0,
		},
		{
			"5 - (3 - 2)",
			4,
		},
		{
			"5 - 3 - 2",
			0,
		},
		{
			"5 - 3 * 2",
			-2,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result, err := Parse(tc.expression)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
