package solution

import (
	"reflect"
	"testing"
)

func TestFindKthNumber(t *testing.T) {
	tests := []struct {
		inputMNK []int
		expected int
	}{
		{[]int{3, 3, 5}, 3},
		{[]int{2, 3, 6}, 6},
		{[]int{9895, 28405, 100787757}, 31666344},
	}
	for i, v := range tests {
		actual := findKthNumber(v.inputMNK[0], v.inputMNK[1], v.inputMNK[2])
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
