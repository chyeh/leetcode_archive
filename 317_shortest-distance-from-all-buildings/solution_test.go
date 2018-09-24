package solution

import (
	"reflect"
	"testing"
)

func TestShortestDistance(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{[][]int{
			{1, 0, 2, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0},
		},
			7,
		},
		{[][]int{
			{1},
		},
			-1,
		},
		{[][]int{
			{0, 2, 1},
			{1, 0, 2},
			{0, 1, 0},
		},
			-1,
		},
		{[][]int{
			{1, 1},
			{0, 1},
		},
			-1,
		},
		{[][]int{
			{1, 1, 1, 1, 1, 0},
			{0, 0, 0, 0, 0, 1},
			{0, 1, 1, 0, 0, 1},
			{1, 0, 0, 1, 0, 1},
			{1, 0, 1, 0, 0, 1},
			{1, 0, 0, 0, 0, 1},
			{0, 1, 1, 1, 1, 0},
		},
			88,
		},
	}
	for i, v := range tests {
		actual := shortestDistance(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
