package solution

import (
	"reflect"
	"testing"
)

func TestValidTree(t *testing.T) {
	tests := []struct {
		inputN     int
		inputEdges [][]int
		expected   bool
	}{
		{
			5,
			[][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 4},
			},
			true,
		},
		{
			5,
			[][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{1, 3},
				{1, 4},
			},
			false,
		},
		{
			2,
			[][]int{
				{1, 0},
			},
			true,
		},
		{
			5,
			[][]int{
				{0, 1},
				{0, 4},
				{1, 4},
				{2, 3},
			},
			false,
		},
	}

	for i, v := range tests {
		actual := validTree(v.inputN, v.inputEdges)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
