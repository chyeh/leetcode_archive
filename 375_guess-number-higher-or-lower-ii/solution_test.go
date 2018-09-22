package solution

import (
	"reflect"
	"testing"
)

func TestGetMoneyAmount(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 6},
		{10, 16},
		{20, 49},
	}
	for i, v := range tests {
		actual := getMoneyAmount(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
