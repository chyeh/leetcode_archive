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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
		return root
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
		return root
	}
	root.Right = insertIntoBST(root.Right, val)
	return root
}
