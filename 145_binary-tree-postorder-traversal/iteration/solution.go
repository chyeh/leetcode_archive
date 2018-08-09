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

func (s *stack) top() *TreeNode {
	return s.data[len(s.data)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	s := NewStack()
	s.push(root)
	prev := &TreeNode{} // dummy node
	for !s.isEmpty() {
		top := s.top()
		if (top.Left == nil && top.Right == nil) || top.Left == prev || top.Right == prev {
			ans = append(ans, top.Val)
			s.pop()
			prev = top
		} else {
			if top.Right != nil {
				s.push(top.Right)
			}
			if top.Left != nil {
				s.push(top.Left)
			}
		}
	}
	return ans
}
