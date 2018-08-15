package solution

import (
	"reflect"
	"testing"
)

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		inputTasks []byte
		inputN     int
		expected   int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E'}, 2, 11},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'C'}, 2, 7},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D'}, 2, 10},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, 2, 12},
	}

	for i, v := range tests {
		actual := leastInterval(v.inputTasks, v.inputN)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
