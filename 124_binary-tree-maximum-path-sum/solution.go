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

func caculateSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	lmax := caculateSum(node.Left, maxSum)
	rmax := caculateSum(node.Right, maxSum)
	sum := max(node.Val, rmax+node.Val)
	sum = max(sum, lmax+node.Val)

	if max(sum, rmax+lmax+node.Val) > *maxSum {
		*maxSum = max(sum, rmax+lmax+node.Val)
	}

	return sum
}

func maxPathSum(root *TreeNode) int {
	maxSum := -2147483648
	caculateSum(root, &maxSum)
	return max(caculateSum(root, &maxSum), maxSum)
}
