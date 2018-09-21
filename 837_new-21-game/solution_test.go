package solution

import (
	"reflect"
	"testing"
)

func TestNew21Game(t *testing.T) {
	tests := []struct {
		inputN   int
		inputK   int
		inputW   int
		expected float64
	}{
		{21, 17, 10, 0.7327777870686082}, // leetcode: 0.73278
		{10, 1, 10, 0.9999999999999999},  // leetcode: 1
		{6, 1, 10, 0.6},
	}
	for i, v := range tests {
		actual := new21Game(v.inputN, v.inputK, v.inputW)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
