package solution

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{
			{-1, 0, 1},
			{-1, -1, 2},
		}},
		{[]int{0, 0, 0}, [][]int{
			{0, 0, 0},
		}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{
			// LeetCode:
			// {-2, 0, 2},
			// {-2, 1, 1},
			{-2, 1, 1},
			{-2, 0, 2},
		}},
		{[]int{0, 0, 0, 0, 0}, [][]int{
			{0, 0, 0},
		}},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, [][]int{
			{0, 0, 0},
		}},
	}
	for i, v := range tests {
		actual := threeSum(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
