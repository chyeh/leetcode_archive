package solution

import (
	"reflect"
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		inputNums []int
		inputK    int
		inputT    int
		expected  bool
	}{
		{[]int{1, 2, 3, 1}, 3, 0, true},
		{[]int{1, 0, 1, 1}, 1, 2, true},
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
		{[]int{2, 2}, 3, 0, true},
		{[]int{10, 15, 18, 24}, 3, 3, true},
		{[]int{}, 0, 0, false},
		{[]int{}, 0, 1, false},
		{[]int{}, 1, 0, false},
		{[]int{}, 1, 1, false},
		{nil, 0, 0, false},
		{nil, 0, 1, false},
		{nil, 1, 0, false},
		{nil, 1, 1, false},
	}
	for i, v := range tests {
		actual := containsNearbyAlmostDuplicate(v.inputNums, v.inputK, v.inputT)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
