package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	var sum int
	if L <= root.Val && root.Val <= R {
		sum += root.Val
	}
	sum += rangeSumBST(root.Left, L, R)
	sum += rangeSumBST(root.Right, L, R)
	return sum
}
