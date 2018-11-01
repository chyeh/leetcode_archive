package solution

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		/*{"babad", "aba"},
		{"cbbd", "bb"},
		{"", ""},*/
		{"ccc", "ccc"},
	}
	for i, v := range tests {
		actual := longestPalindrome(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
