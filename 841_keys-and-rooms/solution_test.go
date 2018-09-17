package solution

import (
	"reflect"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
		{[][]int{{0}}, true},
		{[][]int{{}}, true},
	}
	for i, v := range tests {
		actual := canVisitAllRooms(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
