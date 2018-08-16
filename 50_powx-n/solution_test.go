package solution

import (
	"reflect"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		inputX   float64
		inputN   int
		expected float64
	}{
		{2.0, 10, 1024.0},
		{2.1, 3, 9.261000000000001}, // 9.261
		{2.0, -2, 0.25},
		{2.0, -2, 0.25},
		{0.0, 0, 1.0},
		{0.0, 0, 1.0},
		{1.0, 0, 1.0},
		{-1.0, 0, 1.0},
		{1.0, 2147483647, 1.0},
		{1.0, -2147483648, 1.0},
		{-1.0, 2147483647, -1.0},
		{-1.0, 2147483646, 1.0},
		{-1.0, -2147483648, 1.0},
		{-1.0, -2147483647, -1.0},
	}

	for i, v := range tests {
		actual := myPow(v.inputX, v.inputN)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
