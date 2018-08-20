package solution

import (
	"reflect"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		inputNums   []int
		inputTarget int
		expected    int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{}, 1, 0},
	}

	for i, v := range tests {
		actual := searchInsert(v.inputNums, v.inputTarget)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
