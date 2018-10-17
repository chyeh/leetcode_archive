package solution

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		/*{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},*/
		{"abba", 2},
	}

	for i, v := range tests {
		actual := lengthOfLongestSubstring(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
