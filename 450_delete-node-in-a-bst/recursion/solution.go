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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key == root.Val {
		if root.Right == nil {
			root = root.Left
			return root
		}
		if root.Left == nil {
			root = root.Right
			return root
		}
		rightMinNode := root.Right
		for rightMinNode.Left != nil {
			rightMinNode = rightMinNode.Left
		}
		root.Val = rightMinNode.Val
		root.Right = deleteNode(root.Right, rightMinNode.Val)
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	root.Left = deleteNode(root.Left, key)
	return root
}
