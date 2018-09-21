package solution

import (
	"reflect"
	"testing"
)

func TestNumMatchingSubseq(t *testing.T) {
	tests := []struct {
		inputS     string
		inputWords []string
		expected   int
	}{
		{"abcde", []string{"a", "bb", "acd", "ace"}, 3},
	}
	for i, v := range tests {
		actual := numMatchingSubseq(v.inputS, v.inputWords)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
