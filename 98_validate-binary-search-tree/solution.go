package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val > max || root.Val < min {
		return false
	}
	if root.Left != nil && !isValid(root.Left, min, root.Val-1) {
		return false
	}
	if root.Right != nil && !isValid(root.Right, root.Val+1, max) {
		return false
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, -2147483648, 2147483647)
}
