package solution

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		inputA1  []int
		inputA2  []int
		expected []int
	}{
		{[]int{1, 16, 28, 33, 81, 83}, []int{5, 62, 78, 85}, []int{1, 5, 16, 28, 33, 62, 78, 81, 83, 85}},
		{[]int{}, []int{}, []int{}},
		{nil, nil, nil},
		{nil, []int{}, nil},
		{nil, []int{1}, nil},
		{[]int{}, nil, nil},
		{[]int{1}, nil, nil},
		{[]int{1}, []int{1}, []int{1, 1}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{2}, []int{1}, []int{1, 2}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{1, 3}, []int{2}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 2}, []int{2, 2}, []int{1, 1, 2, 2, 2}},
		{[]int{1, 1, 2}, []int{1, 2}, []int{1, 1, 1, 2, 2}},
	}
	for i, v := range tests {
		actual := merge(v.inputA1, v.inputA2)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		inputA1  []int
		inputA2  []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for i, v := range tests {
		actual := findMedianSortedArrays(v.inputA1, v.inputA2)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
