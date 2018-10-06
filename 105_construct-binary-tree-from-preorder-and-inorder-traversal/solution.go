package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	for i, n := range inorder {
		if n == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:1+i], inorder[:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}
