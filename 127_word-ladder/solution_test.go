package solution

import (
	"reflect"
	"testing"
)

func TestLadderLength(t *testing.T) {
	tests := []struct {
		inputBeginWord string
		inputEndWord   string
		wordList       []string
		expected       int
	}{
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			5,
		},
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			0,
		},
	}

	for i, v := range tests {
		actual := ladderLength(v.inputBeginWord, v.inputEndWord, v.wordList)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
