package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(root *TreeNode, length *int) int {
	if root == nil {
		return 0
	}
	lDepth := depth(root.Left, length)
	rDepth := depth(root.Right, length)
	if lDepth+rDepth > *length {
		*length = lDepth + rDepth
	}
	return max(lDepth, rDepth) + 1
}
func diameterOfBinaryTree(root *TreeNode) int {
	length := 0
	depth(root, &length)
	return length
}
