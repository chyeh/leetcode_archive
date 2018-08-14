package solution

import (
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{2, 2},
		{3, 3},
		{0, 0},
		{1, 1},
	}

	for i, v := range tests {
		actual := climbStairs(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
