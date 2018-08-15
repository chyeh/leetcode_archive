package solution

import (
	"reflect"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"0", 0},
		{"10", 1},
		{"01", 0},
		{"100", 0},
		{"101", 1},
		{"11", 2},
		{"1", 1},
		{"2", 1},
		{"27", 1},
		{"*", 9},
		{"1*", 18},
		{"3*", 9},
		{"*3", 11},
		{"*10*1", 99},
		{"*********", 291868912},
		{"********************", 104671669},
	}

	for i, v := range tests {
		actual := numDecodings(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
