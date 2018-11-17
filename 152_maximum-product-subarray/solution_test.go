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
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{0, 2}, 2},
		{[]int{-2}, -2},
		{[]int{-1, 0, -2, 2}, 2},
		{[]int{-2, 3, -4}, 24},
		{[]int{-4, -3, -2}, 12},
	}

	for i, v := range tests {
		actual := maxProduct(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
