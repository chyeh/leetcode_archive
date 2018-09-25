package solution

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		inputHaystack string
		inputNeedle   string
		expected      int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"", "a", -1},
		{"a", "", 0},
	}

	for i, v := range tests {
		actual := strStr(v.inputHaystack, v.inputNeedle)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
