package solution

import (
	"reflect"
	"testing"
)

func TestFindLengthOfLCIS(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{2, 2, 2, 2, 2}, 1},
		{[]int{1, 2, 3, 2, 3, 4, 5, 6}, 5},
		{[]int{}, 0},
		{[]int{9}, 1},
	}
	for i, v := range tests {
		actual := findLengthOfLCIS(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
