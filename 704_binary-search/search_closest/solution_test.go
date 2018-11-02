package solution

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		inputNums   []int
		inputTarget int
		expected    int
	}{
		{[]int{1, 2, 4, 5, 6, 6, 8, 9}, 11, 9},
		{[]int{2, 5, 6, 7, 8, 8, 9}, 4, 5},
		{[]int{1, 2, 4, 6, 7, 8, 8, 9}, 0, 1},
		{[]int{}, 4, -1},
	}

	for i, v := range tests {
		actual := search(v.inputNums, v.inputTarget)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
