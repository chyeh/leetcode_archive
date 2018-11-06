package solution

import (
	"reflect"
	"testing"
)

func TestCanWin(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1, 2, 3, 1}, 7},
		{[]int{8, 9, 10, 2, 1}, 9},
		{[]int{9, 10, 3, 2, 1}, 11},
		{[]int{2, 2, 1}, 4},
		{[]int{}, 0},
		{nil, 0},
	}

	for i, v := range tests {
		actual := candy(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
