package solution

import (
	"reflect"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"(()", 2},
		{")()())", 4},
		{"(())", 4},
		{")(()())", 6},
		{")(()())(()())", 12},
		{")(()())((()())", 6},
		{"()()))))()()(", 4},
	}

	for i, v := range tests {
		actual := longestValidParentheses(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
