package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestCanMeasureWater(t *testing.T) {
	tests := []struct {
		inputNums []int
		inputK    int
		expected  []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{1, 2}},
	}

	for i, v := range tests {
		actual := topKFrequent(v.inputNums, v.inputK)
		sort.Ints(actual)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
