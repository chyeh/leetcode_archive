package solution

import (
	"reflect"
	"testing"
)

func TestKInversePairs(t *testing.T) {
	tests := []struct {
		inputN   int
		inputK   int
		expected int
	}{
		{1, 0, 1},
		{1, 1, 0},
		{2, 2, 0},
		{3, 0, 1},
		{3, 1, 2},
		{3, 2, 2},
		{10, 0, 1},
		{10, 55, 0},
		{1000, 0, 1},
		{1000, 1, 999},
		{1000, 2, 499499},
		{1000, 1000, 663677020},
	}

	for i, v := range tests {
		actual := kInversePairs(v.inputN, v.inputK)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
