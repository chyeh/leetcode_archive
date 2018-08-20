package solution

import (
	"reflect"
	"testing"
)

type inputStruct struct {
	Ops  []string
	Args [][]int
}

func TestLFUCache(t *testing.T) {
	tests := []struct {
		input    inputStruct
		expected []int
	}{
		{
			inputStruct{
				Ops:  []string{"LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"},
				Args: [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}, {4, 4}, {1}, {3}, {4}},
			},
			[]int{-20000, -20000, -20000, 1, -20000, -1, 3, -20000, -1, 3, 4},
		},
		{
			inputStruct{
				Ops:  []string{"LFUCache", "put", "get"},
				Args: [][]int{{0}, {0, 0}, {0}},
			},
			[]int{-20000, -20000, -1},
		},
		{
			inputStruct{
				Ops:  []string{"LFUCache", "put", "get"},
				Args: [][]int{{2}, {0, 0}, {0}},
			},
			[]int{-20000, -20000, 0},
		},
		{
			inputStruct{
				Ops:  []string{"LFUCache", "get"},
				Args: [][]int{{2}, {0}},
			},
			[]int{-20000, -1},
		},
		{
			inputStruct{
				Ops:  []string{"LFUCache", "put", "put", "get", "get", "get", "put", "put", "get", "get", "get", "get"},
				Args: [][]int{{3}, {2, 2}, {1, 1}, {2}, {1}, {2}, {3, 3}, {4, 4}, {3}, {2}, {1}, {4}},
			},
			[]int{-20000, -20000, -20000, 2, 1, 2, -20000, -20000, -1, 2, 1, 4},
		},
		{
			inputStruct{
				Ops:  []string{"LFUCache", "put", "put", "put", "put", "get", "get", "get", "get", "put", "get", "get", "get", "get", "get"},
				Args: [][]int{{3}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {4}, {3}, {2}, {1}, {5, 5}, {1}, {2}, {3}, {4}, {5}},
			},
			[]int{-20000, -20000, -20000, -20000, -20000, 4, 3, 2, -1, -20000, -1, 2, 3, -1, 5},
		},
		{
			inputStruct{
				Ops:  []string{"LFUCache", "put", "put", "put", "put", "get", "get", "get", "get", "put", "get", "get", "get", "get", "get"},
				Args: [][]int{{3}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {4}, {3}, {2}, {1}, {5, 5}, {1}, {2}, {3}, {4}, {5}},
			},
			[]int{-20000, -20000, -20000, -20000, -20000, 4, 3, 2, -1, -20000, -1, 2, 3, -1, 5},
		},
	}
	var c LFUCache
	for i, v := range tests {
		actual := []int{}
		for ii, in := range v.input.Ops {
			switch in {
			case "LFUCache":
				c = Constructor(v.input.Args[ii][0])
				actual = append(actual, -20000)
			case "put":
				c.Put(v.input.Args[ii][0], v.input.Args[ii][1])
				actual = append(actual, -20000)
			case "get":
				v := c.Get(v.input.Args[ii][0])
				actual = append(actual, v)
			}
		}

		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
