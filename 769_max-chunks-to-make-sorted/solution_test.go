package solution

import (
	"reflect"
	"testing"
)

func TestMaxChunksToSorted(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 3, 2, 1, 0}, 1},
		{[]int{1, 0, 2, 3, 4}, 4},
	}
	for i, v := range tests {
		actual := maxChunksToSorted(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
