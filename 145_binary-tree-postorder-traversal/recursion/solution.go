package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	ans = append(ans, postorderTraversal(root.Left)...)
	ans = append(ans, postorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}
