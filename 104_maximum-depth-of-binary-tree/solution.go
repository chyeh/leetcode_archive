package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution I: Backtracking
/*
func traverse(depth int, root *TreeNode, maxDepth *int) {
	if root == nil {
		return
	}
	depth++
	if depth > *maxDepth {
		*maxDepth = depth
	}
	traverse(depth, root.Left, maxDepth)
	traverse(depth, root.Right, maxDepth)
}

func maxDepth(root *TreeNode) int {
	var maxDepth int
	traverse(0, root, &maxDepth)
	return maxDepth
}
*/

// Solution II: Recursion
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
