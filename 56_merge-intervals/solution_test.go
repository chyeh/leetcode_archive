package solution

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		input    []Interval
		expected []Interval
	}{
		{[]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, []Interval{{1, 6}, {8, 10}, {15, 18}}},
		{[]Interval{{1, 4}, {4, 5}}, []Interval{{1, 5}}},
		{[]Interval{}, []Interval{}},
		{nil, nil},
		{[]Interval{{1, 4}, {1, 5}, {1, 6}, {1, 7}}, []Interval{{1, 7}}},
		{[]Interval{{3, 4}, {2, 5}, {1, 6}, {0, 7}}, []Interval{{0, 7}}},
		{[]Interval{{3, 4}, {2, 5}, {1, 6}, {1, 7}}, []Interval{{1, 7}}},
		{[]Interval{{3, 4}, {2, 4}, {1, 4}, {0, 4}}, []Interval{{0, 4}}},
		{[]Interval{{1, 2}, {2, 4}, {4, 5}, {5, 7}}, []Interval{{1, 7}}},
		{[]Interval{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, []Interval{{1, 2}, {3, 4}, {5, 6}, {7, 8}}},
		{[]Interval{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, []Interval{{1, 10}}},
		{[]Interval{{2, 3}, {4, 5}, {2, 7}, {8, 9}, {9, 10}, {8, 11}}, []Interval{{2, 7}, {8, 11}}},
		{[]Interval{{2, 3}, {4, 5}, {2, 7}, {8, 9}, {9, 10}, {8, 11}, {1, 12}}, []Interval{{1, 12}}},
	}

	for i, v := range tests {
		actual := merge(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
