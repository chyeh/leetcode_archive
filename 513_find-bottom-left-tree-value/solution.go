package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxLevel int
var ans int

func recur(n *TreeNode, level int) {
	if n == nil {
		return
	}
	if n.Left != nil {
		recur(n.Left, level+1)
	}

	if n.Right != nil {
		recur(n.Right, level+1)
	}

	if level > maxLevel {
		maxLevel = level
		ans = n.Val
	}
}

func findBottomLeftValue(root *TreeNode) int {
	recur(root, 1)
	r := ans
	maxLevel = 0
	ans = 0
	return r
}
