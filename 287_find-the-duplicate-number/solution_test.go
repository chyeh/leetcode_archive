package solution

import (
	"reflect"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 1, 2, 3, 5, 2}, 2},
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	}

	for i, v := range tests {
		actual := findDuplicate(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
