package solution

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		inputCandidate []int
		inputTarget    int
		expected       [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, // 1, 1, 2, 5, 6, 7 ,10
			[][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{[]int{2, 5, 2, 1, 2}, 5, // 1, 2, 2, 2, 5
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for i, v := range tests {
		actual := combinationSum2(v.inputCandidate, v.inputTarget)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
