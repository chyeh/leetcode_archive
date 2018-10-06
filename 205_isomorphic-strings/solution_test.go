package solution

import (
	"reflect"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		inputS   string
		inputT   string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "aa", false},
	}

	for i, v := range tests {
		actual := isIsomorphic(v.inputS, v.inputT)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
