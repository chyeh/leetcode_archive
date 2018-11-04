package solution

import (
	"reflect"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
	}

	for i, v := range tests {
		actual := lengthOfLIS(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
