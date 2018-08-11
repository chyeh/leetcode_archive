package solution

import (
	"reflect"
	"testing"
)

func TestCrackSafe(t *testing.T) {
	tests := []struct {
		inputN   int
		inputK   int
		expected string
	}{
		{1, 1, "0"},
		{1, 2, "01"},
		{2, 2, "00110"},
		{1, 3, "012"},                        //leetcode: 210
		{2, 5, "00102030411213142232433440"}, // leetcode: 04433423224131211403020100
	}

	for i, v := range tests {
		actual := crackSafe(v.inputN, v.inputK)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
