package solution

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		input    int
		expected [][]string
	}{
		{4, [][]string{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		},
		},
	}

	for i, v := range tests {
		actual := solveNQueens(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
