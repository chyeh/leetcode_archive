package solution

import "container/list"

type queue struct {
	*list.List
}

func NewQueue() *queue {
	return &queue{
		list.New(),
	}
}

func (q *queue) push(d interface{}) {
	q.PushBack(d)
}

func (q *queue) pop() interface{} {
	return q.Remove(q.Front())
}

func (q *queue) isEmpty() bool {
	return q.Len() == 0
}

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

type TreeNodeWithLevel struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := NewQueue()
	q.push(TreeNodeWithLevel{root, 0})

	for !q.isEmpty() {
		popped := q.pop().(TreeNodeWithLevel)
		if popped.level == len(ans) {
			ans = append(ans, []int{})
		}
		ans[popped.level] = append(ans[popped.level], popped.node.Val)

		if popped.node.Left != nil {
			q.push(TreeNodeWithLevel{popped.node.Left, popped.level + 1})
		}
		if popped.node.Right != nil {
			q.push(TreeNodeWithLevel{popped.node.Right, popped.level + 1})
		}
	}

	return ans
}
