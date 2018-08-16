package solution

import (
	"reflect"
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		inputS   string
		inputT   string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "aa", ""},
		{"a", "b", ""},
		{"ab", "b", "b"},
		{"ab", "a", "a"},
		{"bba", "ab", "ba"},
		{"aa", "aa", "aa"},
		{"aaa", "aa", "aa"},
		{"abc", "aa", ""},
	}

	for i, v := range tests {
		actual := minWindow(v.inputS, v.inputT)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
