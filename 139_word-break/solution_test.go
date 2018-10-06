package solution

import (
	"reflect"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		inputS        string
		inputWordDict []string
		expected      bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for i, v := range tests {
		actual := wordBreak(v.inputS, v.inputWordDict)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
