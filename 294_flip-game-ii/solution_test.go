package solution

import (
	"reflect"
	"testing"
)

func TestCanWin(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"++++", true},
		{"++++++", true},
	}

	for i, v := range tests {
		actual := canWin(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
