package solution

import (
	"reflect"
	"testing"
)

func TestFindLUSlength(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"aba", "cdc", "eae"}, 3},
		{[]string{"aaa", "aa", "a"}, 3},
	}
	for i, v := range tests {
		actual := findLUSlength(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
