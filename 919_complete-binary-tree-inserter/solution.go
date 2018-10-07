package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(root *TreeNode, v int) int {
	var p *TreeNode
	curr := root
	for {
		p = curr
		if isFull(curr) || !isFull(curr.Left) {
			curr = curr.Left
			if curr == nil {
				p.Left = &TreeNode{Val: v}
				break
			}
		} else {
			curr = curr.Right
			if curr == nil {
				p.Right = &TreeNode{Val: v}
				break
			}
		}

	}
	return p.Val
}

func isFull(node *TreeNode) bool {
	if node == nil {
		return true
	}
	l, r := node.Left, node.Right
	for l != nil && r != nil {
		l = l.Left
		r = r.Right
	}
	return l == r
}

type CBTInserter struct {
	root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{
		root: root,
	}
}

func (this *CBTInserter) Insert(v int) int {
	return insert(this.root, v)
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
