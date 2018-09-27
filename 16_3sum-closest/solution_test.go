package solution

import (
	"reflect"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		inputNums   []int
		inputTarget int
		expected    int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 1, 2}, 3, 3},
	}
	for i, v := range tests {
		actual := threeSumClosest(v.inputNums, v.inputTarget)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
