package solution

import (
	"reflect"
	"testing"
)

func TestCountBattleships(t *testing.T) {
	tests := []struct {
		input    [][]byte
		expected int
	}{
		{[][]byte{
			{'X', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
			{'.', '.', '.', 'X'}}, 2},
		{[][]byte{
			{'X', 'X', 'X'}}, 1},
		{[][]byte{
			{'X'},
			{'X'},
			{'X'}}, 1},
	}
	for i, v := range tests {
		actual := countBattleships(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
