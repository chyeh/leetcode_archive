package solution

import (
	"reflect"
	"testing"
)

func TestFindLUSlength(t *testing.T) {
	tests := []struct {
		inputA   string
		inputB   string
		expected int
	}{
		{"aba", "cdc", 3},
	}
	for i, v := range tests {
		actual := findLUSlength(v.inputA, v.inputB)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
