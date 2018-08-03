package solution

import (
	"reflect"
	"testing"
)

func TestExclusiveTime(t *testing.T) {
	tests := []struct {
		inputN    int
		inputLogs []string
		expected  []int
	}{
		{
			2, []string{
				"0:start:0",
				"1:start:2",
				"1:end:5",
				"0:end:6",
			}, []int{3, 4},
		},
	}

	for i, v := range tests {
		actual := exclusiveTime(v.inputN, v.inputLogs)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
