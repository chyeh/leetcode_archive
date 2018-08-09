package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	data []*TreeNode
}

func NewStack() *stack {
	return &stack{
		data: make([]*TreeNode, 0),
	}
}

func (s *stack) push(d *TreeNode) {
	s.data = append(s.data, d)
}

func (s *stack) pop() *TreeNode {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	s := NewStack()
	for root != nil || !s.isEmpty() {
		if root != nil {
			ans = append(ans, root.Val)
			s.push(root.Right)
			root = root.Left
		} else {
			root = s.pop()
		}
	}
	return ans
}
