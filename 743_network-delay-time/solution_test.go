package solution

import (
	"reflect"
	"testing"
)

func TestCrackSafe(t *testing.T) {
	tests := []struct {
		inputTimes [][]int
		inputN     int
		inputK     int
		expected   int
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{[][]int{{1, 2, 1}}, 2, 2, -1},
		{[][]int{{4, 2, 76}, {1, 3, 79}, {3, 1, 81}, {4, 3, 30}, {2, 1, 47}, {1, 5, 61}, {1, 4, 99}, {3, 4, 68}, {3, 5, 46}, {4, 1, 6}, {5, 4, 7}, {5, 3, 44}, {4, 5, 19}, {2, 3, 13}, {3, 2, 18}, {1, 2, 0}, {5, 1, 25}, {2, 5, 58}, {2, 4, 77}, {5, 2, 74}}, 5, 3, 59},
		{[][]int{{2, 4, 10}, {5, 2, 38}, {3, 4, 33}, {4, 2, 76}, {3, 2, 64}, {1, 5, 54}, {1, 4, 98}, {2, 3, 61}, {2, 1, 0}, {3, 5, 77}, {5, 1, 34}, {3, 1, 79}, {5, 3, 2}, {1, 2, 59}, {4, 3, 46}, {5, 4, 44}, {2, 5, 89}, {4, 5, 21}, {1, 3, 86}, {4, 1, 95}}, 5, 1, 69},
	}

	for i, v := range tests {
		actual := networkDelayTime(v.inputTimes, v.inputN, v.inputK)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
