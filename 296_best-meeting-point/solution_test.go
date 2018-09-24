package solution

import (
	"reflect"
	"testing"
)

func TestMinTotalDistance(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{[][]int{
			{1, 0, 0, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0},
		},
			6,
		},
		{[][]int{
			{1, 1},
		},
			1,
		},
		{[][]int{
			{0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0},
		},
			11,
		},
	}
	for i, v := range tests {
		actual := minTotalDistance(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
