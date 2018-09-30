package solution

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		inputNums []int
		inputK    int
		expected  []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{}, 0, []int{}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{7, 2, 4}, 2, []int{7, 4}},
	}

	for i, v := range tests {
		actual := maxSlidingWindow(v.inputNums, v.inputK)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
