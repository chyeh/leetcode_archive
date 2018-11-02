package solution

import (
	"reflect"
	"testing"
)

func TestCanWin(t *testing.T) {
	tests := []struct {
		inputMaxChoosableInteger int
		inputDesiredTotal        int
		expected                 bool
	}{
		{10, 11, false},
		{10, 40, false},
		{20, 210, false},
	}

	for i, v := range tests {
		actual := canIWin(v.inputMaxChoosableInteger, v.inputDesiredTotal)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
