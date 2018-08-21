package solution

import (
	"reflect"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		inputNums []int
		inputK    int
		expected  bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{}, 0, false},
		{[]int{}, 1, false},
		{nil, 0, false},
		{nil, 1, false},
	}
	for i, v := range tests {
		actual := containsNearbyDuplicate(v.inputNums, v.inputK)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
