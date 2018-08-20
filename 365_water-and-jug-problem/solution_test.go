package solution

import (
	"reflect"
	"testing"
)

func TestCanMeasureWater(t *testing.T) {
	tests := []struct {
		inputX   int
		inputY   int
		inputZ   int
		expected bool
	}{
		{3, 5, 4, true},
		{2, 6, 5, false},
		{0, 0, 0, true},
		{2002, 847, 77, true},
		{847, 2002, 77, true},
		{97, 11, 1, true},
	}

	for i, v := range tests {
		actual := canMeasureWater(v.inputX, v.inputY, v.inputZ)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
