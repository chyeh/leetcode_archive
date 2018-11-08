package solution

import (
	"reflect"
	"testing"
)

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		inputN       int
		inputFlights [][]int
		inputSrc     int
		inputDst     int
		inputK       int
		expected     int
	}{
		{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1, 200},
		{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0, 500},
		{4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1, 6},
		{7, [][]int{{0, 3, 7}, {4, 5, 3}, {6, 4, 8}, {2, 0, 10}, {6, 5, 6}, {1, 2, 2}, {2, 5, 9}, {2, 6, 8}, {3, 6, 3}, {4, 0, 10}, {4, 6, 8}, {5, 2, 6}, {1, 4, 3}, {4, 1, 6}, {0, 5, 10}, {3, 1, 5}, {4, 3, 1}, {5, 4, 10}, {0, 1, 6}}, 2, 4, 1, 16},
	}

	for i, v := range tests {
		actual := findCheapestPrice(v.inputN, v.inputFlights, v.inputSrc, v.inputDst, v.inputK)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
