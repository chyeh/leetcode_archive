package solution

import (
	"reflect"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		inputS   string
		inputT   string
		expected bool
	}{
		{"rat", "car", false},
		{"anagram", "nagaram", true},
	}
	for i, v := range tests {
		actual := isAnagram(v.inputS, v.inputT)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
