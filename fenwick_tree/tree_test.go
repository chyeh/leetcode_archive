package fenwick_tree

import (
	"reflect"
	"testing"
)

func TestLSB(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{10, 2},
		{6, 2},
		{11, 1},
		{8, 8},
	}
	for i, v := range tests {
		actual := lsb(v.input)
		if actual != v.expected {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 2, 3, 1, 7, 2, 5, 4, 21, 6, 13, 8, 30},
		},
	}
	for i, v := range tests {
		actual := ([]int)(*New(v.input))
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		inputArray []int
		inputIndex int
		expected   int
	}{
		{
			[]int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 0,
		},
		{
			[]int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}, 1, 2,
		},
		{
			[]int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}, 2, 3,
		},
		{
			[]int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}, 7, 16,
		},
		{
			[]int{1, 3, 5}, 3, 9,
		},
	}
	for i, v := range tests {
		actual := New(v.inputArray).sum(v.inputIndex)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		inputArray   []int
		inputAddArgs []struct {
			index  int
			number int
		}
		expectedArray [][]int
	}{
		{
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]struct {
				index  int
				number int
			}{{1, 2}, {2, 1}, {3, 1}},
			[][]int{
				[]int{0, 2, 2, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0},
				[]int{0, 2, 3, 0, 3, 0, 0, 0, 3, 0, 0, 0, 0},
				[]int{0, 2, 3, 1, 4, 0, 0, 0, 4, 0, 0, 0, 0},
			},
		},
	}
	for i, v := range tests {
		tree := New(v.inputArray)
		for ii, arg := range v.inputAddArgs {
			tree.add(arg.index, arg.number)
			actual := ([]int)(*tree)
			expected := v.expectedArray[ii]
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Error on case %d-%d: %v(actual) != %v(expected)", i, ii, actual, expected)
			}
			t.Logf("Case %d-%d: %v(actual) == %v(expected)", i, ii, actual, expected)
		}
	}
}
