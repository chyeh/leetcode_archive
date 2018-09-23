package solution

import (
	"reflect"
	"testing"
)

func TestPredictTheWinner(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 5, 2}, false},
		{[]int{1, 5, 233, 7}, true},
	}
	for i, v := range tests {
		actual := PredictTheWinner(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
