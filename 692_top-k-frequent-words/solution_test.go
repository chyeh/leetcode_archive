package solution

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		inputWords []string
		inputK     int
		expected   []string
	}{
		{
			[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2,
			[]string{"i", "love"},
		},
		{
			[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4,
			[]string{"the", "is", "sunny", "day"},
		},
		{
			[]string{"i", "love", "leetcode", "i", "love", "coding"}, 3,
			[]string{"i", "love", "coding"},
		},
		{
			[]string{"aaa", "aa", "a"}, 1,
			[]string{"a"},
		},
	}
	for i, v := range tests {
		actual := topKFrequent(v.inputWords, v.inputK)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
