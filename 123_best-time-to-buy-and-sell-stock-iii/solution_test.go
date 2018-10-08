package solution

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, 13},
		{[]int{2, 1, 2, 0, 1}, 2},
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{2}, 0},
	}

	for i, v := range tests {
		actual := maxProfit(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
