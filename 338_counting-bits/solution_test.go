package solution

import (
	"reflect"
	"testing"
)

func TestCountBattleships(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for i, v := range tests {
		actual := countBits(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
