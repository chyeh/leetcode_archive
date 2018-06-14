package solution

import (
	"reflect"
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected int
	}{
		{&TreeNode{
			Val: 0,
		}, 0},
		{&TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3},
		}, 1},
		{&TreeNode{
			Val:  1,
			Left: nil, Right: &TreeNode{Val: 1},
		}, 1},
		{&TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 4},
			}, Right: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 5,
					Left: &TreeNode{Val: 7},
				}, Right: &TreeNode{Val: 6},
			},
		}, 7},
	}
	for i, v := range tests {
		actual := findBottomLeftValue(v.input)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}
