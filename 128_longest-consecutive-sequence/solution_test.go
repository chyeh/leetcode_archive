package solution

import (
	"reflect"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{1, 2, 0, 1}, 3},
	}

	for i, v := range tests {
		actual := longestConsecutive(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
