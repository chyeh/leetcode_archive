package solution

import (
	"reflect"
	"testing"
)

type inputStruct struct {
	Ops  []string
	Args [][]int
}

func TestLRUCache(t *testing.T) {
	tests := []struct {
		input    inputStruct
		expected []int
	}{
		{
			inputStruct{
				Ops:  []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				Args: [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			},
			[]int{-20000, -20000, -20000, 1, -20000, -1, -20000, -1, 3, 4},
		},
		{
			inputStruct{
				Ops:  []string{"LRUCache", "put", "get"},
				Args: [][]int{{0}, {0, 0}, {0}},
			},
			[]int{-20000, -20000, -1},
		},
		{
			inputStruct{
				Ops:  []string{"LRUCache", "put", "get"},
				Args: [][]int{{1}, {2, 1}, {2}},
			},
			[]int{-20000, -20000, 1},
		},
		{
			inputStruct{
				Ops:  []string{"LRUCache", "put", "put", "get", "put", "put", "get"},
				Args: [][]int{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2}},
			},
			[]int{-20000, -20000, -20000, 2, -20000, -20000, -1},
		},
		{
			inputStruct{
				Ops:  []string{"LRUCache", "put", "get", "put", "get", "get"},
				Args: [][]int{{1}, {2, 1}, {2}, {3, 2}, {2}, {3}},
			},
			[]int{-20000, -20000, 1, -20000, -1, 2},
		},
		{
			inputStruct{
				Ops:  []string{"LRUCache", "put", "put", "put", "put", "get", "get", "get", "get", "put", "get", "get", "get", "get", "get"},
				Args: [][]int{{3}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {4}, {3}, {2}, {1}, {5, 5}, {1}, {2}, {3}, {4}, {5}},
			},
			[]int{-20000, -20000, -20000, -20000, -20000, 4, 3, 2, -1, -20000, -1, 2, 3, -1, 5},
		},
		{
			inputStruct{
				Ops: []string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put",
					"get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put",
					"put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put",
					"put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get",
					"put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"},
				Args: [][]int{{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12},
					{7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24},
					{5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10},
					{4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20},
					{11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17},
					{9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}},
			},
			[]int{-20000, -20000, -20000, -20000, -20000, -20000, -1, -20000, 19, 17, -20000, -1, -20000, -20000, -20000, -1, -20000, -1, 5, -1, 12, -20000, -20000, 3, 5, 5, -20000, -20000, 1, -20000, -1, -20000, 30, 5, 30,
				-20000, -20000, -20000, -1, -20000, -1, 24, -20000, -20000, 18, -20000, -20000, -20000, -20000, -1, -20000, -20000, 18, -20000, -20000, -1, -20000, -20000, -20000, -20000, -20000, 18, -20000, -20000, -1, -20000, 4, 29, 30, -20000,
				12, -1, -20000, -20000, -20000, -20000, 29, -20000, -20000, -20000, -20000, 17, 22, 18, -20000, -20000, -20000, -1, -20000, -20000, -20000, 20, -20000, -20000, -20000, -1, 18, 18, -20000, -20000, -20000, -20000, 20, -20000, -20000,
				-20000, -20000, -20000, -20000, -20000},
		},
	}
	var c LRUCache
	for i, v := range tests {
		actual := []int{}
		for ii, in := range v.input.Ops {
			switch in {
			case "LRUCache":
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
