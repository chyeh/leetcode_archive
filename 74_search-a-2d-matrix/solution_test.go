package solution

import (
	"reflect"
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		inputMatrix [][]int
		inputTarget int
		expected    bool
	}{
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			3, true},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			13, false},
		{
			[][]int{},
			1, false},
	}

	for i, v := range tests {
		actual := searchMatrix(v.inputMatrix, v.inputTarget)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
