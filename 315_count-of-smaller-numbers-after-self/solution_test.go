package solution

import (
	"reflect"
	"testing"
)

func TestCountSmaller(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{5, 5, 2, 6, 1}, []int{2, 2, 1, 1, 0}},
		{[]int{5, 2, 2, 6, 1}, []int{3, 1, 1, 1, 0}},
		{[]int{5, 2, 6, 6, 1}, []int{2, 1, 1, 1, 0}},
		{[]int{2, 0, 1}, []int{2, 0, 0}},
		{
			[]int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41},
			[]int{10, 27, 10, 35, 12, 22, 28, 8, 19, 2, 12, 2, 9, 6, 12, 5, 17, 9, 19, 12, 14, 6, 12, 5, 12, 3, 0, 10, 0, 7, 8, 4, 0, 0, 4, 3, 2, 0, 1, 0},
		},
		{[]int{}, []int{}},
		{nil, []int{}},
		{[]int{1}, []int{0}},
	}
	for i, v := range tests {
		actual := countSmaller(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
