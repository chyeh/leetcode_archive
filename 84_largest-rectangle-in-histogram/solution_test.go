package solution

import (
	"reflect"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		//{[]int{0, 0, 0, 0, 0, 0, 0, 0, 2147483647}, 2147483647},
	}

	for i, v := range tests {
		actual := largestRectangleArea(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
