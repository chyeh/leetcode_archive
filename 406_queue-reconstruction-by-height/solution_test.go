package solution

import (
	"reflect"
	"testing"
)

func TestInsertPerson(t *testing.T) {
	tests := []struct {
		inputArray    [][]int
		inputI        int
		expectedArray [][]int
	}{
		{
			[][]int{{7, 0}, {7, 1}, {6, 1}, {5, 0}, {5, 2}, {4, 4}},
			2,
			[][]int{{7, 0}, {6, 1}, {7, 1}, {5, 0}, {5, 2}, {4, 4}},
		},
		{
			[][]int{{7, 0}, {6, 1}, {7, 1}, {5, 0}, {5, 2}, {4, 4}},
			3,
			[][]int{{5, 0}, {7, 0}, {6, 1}, {7, 1}, {5, 2}, {4, 4}},
		},
	}
	for i, v := range tests {
		actual := insertPerson(v.inputArray, v.inputI)
		if !reflect.DeepEqual(actual, v.expectedArray) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expectedArray)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expectedArray)
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		inputArray [][]int
		inputX     int
		inputY     int
		expected   [][]int
	}{
		{
			[][]int{{1, 2}, {1, 3}},
			0,
			1,
			[][]int{{1, 3}, {1, 2}},
		},
	}
	for i, v := range tests {
		actual := swap(v.inputArray, v.inputX, v.inputY)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

func TestReconstructQueue(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{ // test cast 1:
			[][]int{
				{7, 0},
				{4, 4},
				{7, 1},
				{5, 0},
				{6, 1},
				{5, 2},
			},
			[][]int{
				{5, 0},
				{7, 0},
				{5, 2},
				{6, 1},
				{4, 4},
				{7, 1},
			},
		}, // :~)
		{ // test cast 2:
			[][]int{
				{7, 0},
				{4, 4},
				{7, 1},
				{5, 0},
				{6, 1},
				{5, 2},
			},
			[][]int{
				{5, 0},
				{7, 0},
				{5, 2},
				{6, 1},
				{4, 4},
				{7, 1},
			},
		}, // :~)
		{
			[][]int{
				{9, 0},
				{7, 0},
				{1, 9},
				{3, 0},
				{2, 7},
				{5, 3},
				{6, 0},
				{3, 4},
				{6, 2},
				{5, 2},
			},
			[][]int{
				{3, 0},
				{6, 0},
				{7, 0},
				{5, 2},
				{3, 4},
				{5, 3},
				{6, 2},
				{2, 7},
				{9, 0},
				{1, 9},
			},
		}, // :~)
	}
	for i, v := range tests {
		actual := reconstructQueue(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
