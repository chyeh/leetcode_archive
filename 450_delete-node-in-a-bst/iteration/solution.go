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

func delete(node *TreeNode) *TreeNode {
	if node.Right == nil {
		node = node.Left
		return node
	}
	if node.Left == nil {
		node = node.Right
		return node
	}

	p := node
	i := node.Right
	for i.Left != nil {
		p = i
		i = i.Left
	}
	node.Val = i.Val
	if p == node {
		node.Right = i.Right
	} else {
		p.Left = i.Right
	}
	return node
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	i := root
	var p *TreeNode
	for i != nil {
		if key == i.Val {
			break
		}
		p = i
		if key > i.Val {
			i = i.Right
		} else {
			i = i.Left
		}
	}
	if i == nil {
		return root
	}
	if p == nil {
		return delete(root)
	}

	if p.Left == i {
		p.Left = delete(i)
	} else {
		p.Right = delete(i)
	}
	return root
}
