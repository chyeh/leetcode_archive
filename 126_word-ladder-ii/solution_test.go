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
		expected       [][]string
	}{
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			[][]string{
				[]string{"hit", "hot", "dot", "dog", "cog"},
				[]string{"hit", "hot", "lot", "log", "cog"},
			},
		},
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			nil,
		},
		{
			"hot",
			"dog",
			[]string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"},
			[][]string{
				[]string{"hot", "hog", "dog"},
				[]string{"hot", "dot", "dog"},
				// LeedCode:
				// []string{"hot", "dot", "dog"},
				// []string{"hot", "hog", "dog"},
			},
		},
		{
			"red",
			"tax",
			[]string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
			[][]string{
				[]string{"red", "ted", "tex", "tax"},
				[]string{"red", "rex", "tex", "tax"},
				[]string{"red", "ted", "tad", "tax"},
				// LeetCode:
				// []string{"red", "ted", "tad", "tax"},
				// []string{"red", "ted", "tex", "tax"},
				// []string{"red", "rex", "tex", "tax"},
			},
		},
		{
			"qa",
			"sq",
			[]string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"},
			[][]string{
				{"qa", "ca", "ci", "si", "sq"}, {"qa", "ta", "ti", "si", "sq"}, {"qa", "la", "li", "si", "sq"}, {"qa", "ba", "bi", "si", "sq"}, {"qa", "ha", "hi", "si", "sq"}, {"qa", "pa", "pi", "si", "sq"}, {"qa", "ma", "mi", "si", "sq"}, {"qa", "na", "ni", "si", "sq"}, {"qa", "fa", "fe", "se", "sq"}, {"qa", "ra", "re", "se", "sq"}, {"qa", "ga", "ge", "se", "sq"}, {"qa", "ya", "ye", "se", "sq"}, {"qa", "ca", "co", "so", "sq"}, {"qa", "ca", "cm", "sm", "sq"}, {"qa", "ca", "cr", "sr", "sq"}, {"qa", "ta", "to", "so", "sq"}, {"qa", "ta", "tb", "sb", "sq"}, {"qa", "ta", "tm", "sm", "sq"}, {"qa", "ta", "th", "sh", "sq"}, {"qa", "ta", "tc", "sc", "sq"}, {"qa", "la", "le", "se", "sq"}, {"qa", "la", "lo", "so", "sq"}, {"qa", "la", "ln", "sn", "sq"}, {"qa", "la", "lr", "sr", "sq"}, {"qa", "la", "lt", "st", "sq"}, {"qa", "ba", "be", "se", "sq"}, {"qa", "ba", "br", "sr", "sq"}, {"qa", "ha", "he", "se", "sq"}, {"qa", "ha", "ho", "so", "sq"}, {"qa", "pa", "po", "so", "sq"}, {"qa", "pa", "pb", "sb", "sq"}, {"qa", "pa", "pm", "sm", "sq"}, {"qa", "pa", "ph", "sh", "sq"}, {"qa", "pa", "pt", "st", "sq"}, {"qa", "ma", "me", "se", "sq"}, {"qa", "ma", "mo", "so", "sq"}, {"qa", "ma", "mb", "sb", "sq"}, {"qa", "ma", "mn", "sn", "sq"}, {"qa", "ma", "mr", "sr", "sq"}, {"qa", "ma", "mt", "st", "sq"}, {"qa", "na", "ne", "se", "sq"}, {"qa", "na", "no", "so", "sq"}, {"qa", "na", "nb", "sb", "sq"}, {"qa", "fa", "fm", "sm", "sq"}, {"qa", "fa", "fr", "sr", "sq"}, {"qa", "ra", "rb", "sb", "sq"}, {"qa", "ra", "rn", "sn", "sq"}, {"qa", "ra", "rh", "sh", "sq"}, {"qa", "ga", "go", "so", "sq"}, {"qa", "ya", "yo", "so", "sq"}, {"qa", "ya", "yb", "sb", "sq"},
			},
			// LeetCode:
			// [["qa","pa","pt","st","sq"],["qa","la","lt","st","sq"],["qa","ma","mt","st","sq"],["qa","ca","cr","sr","sq"],["qa","la","lr","sr","sq"],["qa","fa","fr","sr","sq"],["qa","ba","br","sr","sq"],["qa","ma","mr","sr","sq"],["qa","ca","ci","si","sq"],["qa","na","ni","si","sq"],["qa","la","li","si","sq"],["qa","ta","ti","si","sq"],["qa","pa","pi","si","sq"],["qa","ba","bi","si","sq"],["qa","ha","hi","si","sq"],["qa","ma","mi","si","sq"],["qa","pa","ph","sh","sq"],["qa","ra","rh","sh","sq"],["qa","ta","th","sh","sq"],["qa","ca","co","so","sq"],["qa","ga","go","so","sq"],["qa","ta","to","so","sq"],["qa","na","no","so","sq"],["qa","la","lo","so","sq"],["qa","pa","po","so","sq"],["qa","ya","yo","so","sq"],["qa","ma","mo","so","sq"],["qa","ha","ho","so","sq"],["qa","la","ln","sn","sq"],["qa","ra","rn","sn","sq"],["qa","ma","mn","sn","sq"],["qa","ca","cm","sm","sq"],["qa","ta","tm","sm","sq"],["qa","pa","pm","sm","sq"],["qa","fa","fm","sm","sq"],["qa","ta","tc","sc","sq"],["qa","na","nb","sb","sq"],["qa","pa","pb","sb","sq"],["qa","ra","rb","sb","sq"],["qa","ya","yb","sb","sq"],["qa","ma","mb","sb","sq"],["qa","ta","tb","sb","sq"],["qa","ga","ge","se","sq"],["qa","la","le","se","sq"],["qa","na","ne","se","sq"],["qa","ra","re","se","sq"],["qa","ba","be","se","sq"],["qa","ya","ye","se","sq"],["qa","fa","fe","se","sq"],["qa","ha","he","se","sq"],["qa","ma","me","se","sq"]]

		},
	}

	for i, v := range tests {
		actual := findLadders(v.inputBeginWord, v.inputEndWord, v.wordList)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
