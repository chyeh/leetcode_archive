package solution

import (
	"reflect"
	"testing"
)

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		input    []Interval
		expected int
	}{
		{[]Interval{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[]Interval{{7, 10}, {2, 4}}, 1},
	}
	for i, v := range tests {
		actual := minMeetingRooms(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
