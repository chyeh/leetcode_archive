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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	var p *TreeNode
	curr := root.Left
	for curr != nil {
		p = curr
		curr = curr.Right
	}
	if p == nil {
		return
	}
	p.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}
