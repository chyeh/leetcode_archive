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
	curr := root
	var prev *TreeNode
	for curr != nil {
		prev = curr
		if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	if val < prev.Val {
		prev.Left = &TreeNode{
			Val: val,
		}
	} else {
		prev.Right = &TreeNode{
			Val: val,
		}
	}
	return root
}
