package solution

import (
	"testing"
)

func TestIsArithmatic(t *testing.T) {
	tests := []struct {
		inputA   []int
		inputI   int
		inputJ   int
		expected bool
	}{
		{[]int{1, 1, 1}, 0, 2, true},
		{[]int{1, 3, 5}, 0, 2, true},
		{[]int{1, 3, 6}, 0, 2, false},
	}
	for i, v := range tests {
		actual := isArithmatic(v.inputA, v.inputI, v.inputJ)
		if actual != v.expected {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 1, 1}, 1},
	}
	for i, v := range tests {
		actual := numberOfArithmeticSlices(v.input)
		if actual != v.expected {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
