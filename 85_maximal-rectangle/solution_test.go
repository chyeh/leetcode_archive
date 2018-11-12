package solution

import (
	"reflect"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		input    [][]byte
		expected int
	}{
		/*{[][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}, 6},*/
		{[][]byte{
			{'0', '1', '1', '0', '1'},
			{'1', '1', '0', '1', '0'},
			{'0', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '1'},
			{'0', '0', '0', '0', '0'},
		}, 9},
	}

	for i, v := range tests {
		actual := maximalRectangle(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
