package solution

import (
	"reflect"
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		inputS   string
		inputP   string
		expected bool
	}{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{"adbeb", "*a*b", true},
	}

	for i, v := range tests {
		actual := isMatch(v.inputS, v.inputP)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
