package solution

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		inputNums []int
		inputK    int
		expected  int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{3, 4, 7, 2, -3, 1, 4, 2}, 7, 4},
		{[]int{1}, 0, 0},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 55},
	}

	for i, v := range tests {
		actual := subarraySum(v.inputNums, v.inputK)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
