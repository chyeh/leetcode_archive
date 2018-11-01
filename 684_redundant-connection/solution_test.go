package solution

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			[][]int{{1, 2}, {1, 3}, {2, 3}},
			[]int{2, 3},
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			[]int{1, 4},
		},
		{
			[][]int{{9, 10}, {5, 8}, {2, 6}, {1, 5}, {3, 8}, {4, 9}, {8, 10}, {4, 10}, {6, 8}, {7, 9}},
			[]int{4, 10},
		},
		{
			[][]int{{3, 4}, {1, 2}, {2, 4}, {3, 5}, {2, 5}},
			[]int{2, 5},
		},
		{
			[][]int{{1, 5}, {3, 4}, {3, 5}, {4, 5}, {2, 4}},
			[]int{4, 5},
		},
	}

	for i, v := range tests {
		actual := findRedundantConnection(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
