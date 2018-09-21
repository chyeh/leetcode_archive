package solution

import (
	"reflect"
	"testing"
)

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		input    []Interval
		expected bool
	}{
		{[]Interval{{0, 30}, {5, 10}, {15, 20}}, false},
		{[]Interval{{7, 10}, {2, 4}}, true},
		{[]Interval{{9, 10}, {4, 9}, {4, 17}}, false},
		{[]Interval{{13, 5}, {1, 13}}, true},
		{[]Interval{{16, 22}, {28, 45}, {3, 9}, {46, 50}, {13, 14}}, true},
	}
	for i, v := range tests {
		actual := canAttendMeetings(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
