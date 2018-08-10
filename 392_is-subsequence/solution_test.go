package solution

import (
	"reflect"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		inputS   string
		inputT   string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"ace", "abcde", true},
		{"aec", "abcde", false},
	}

	for i, v := range tests {
		actual := isSubsequence(v.inputS, v.inputT)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
