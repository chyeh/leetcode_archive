package solution

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		inputArr []int
		inputK   int
		inputX   int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		{[]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}, 9, 4, []int{0, 1, 1, 1, 2, 3, 6, 7, 8}},
		{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5, []int{3, 3, 4}},
		{[]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8}, 2, 2, []int{1, 3}},
		{[]int{1, 2, 3, 3, 6, 6, 7, 7, 9, 9}, 8, 8, []int{3, 3, 6, 6, 7, 7, 9, 9}},
	}

	for i, v := range tests {
		actual := findClosestElements(v.inputArr, v.inputK, v.inputX)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
