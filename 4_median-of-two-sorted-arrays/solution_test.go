package solution

import (
	"reflect"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		inputA1  []int
		inputA2  []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{3, 4}, 3.5},
	}
	for i, v := range tests {
		actual := findMedianSortedArrays(v.inputA1, v.inputA2)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
