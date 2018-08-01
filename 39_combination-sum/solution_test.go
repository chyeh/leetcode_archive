package solution

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		inputCandidate []int
		inputTarget    int
		expected       [][]int
	}{
		{[]int{2, 3, 6, 7}, 7,
			[][]int{
				{2, 2, 3},
				{7},
			},
		},
		{[]int{2, 3, 5}, 8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
	}

	for i, v := range tests {
		actual := combinationSum(v.inputCandidate, v.inputTarget)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
