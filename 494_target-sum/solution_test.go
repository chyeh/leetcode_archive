package solution

import (
	"reflect"
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		inputNums []int
		inputS    int
		expected  int
	}{
		{
			[]int{1, 1, 1, 1, 1}, 3, 5,
		},
		{
			[]int{1, 0}, 1, 2,
		},
	}

	for i, v := range tests {
		actual := findTargetSumWays(v.inputNums, v.inputS)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
