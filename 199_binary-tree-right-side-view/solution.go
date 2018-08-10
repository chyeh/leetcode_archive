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

type item struct {
	*TreeNode
	level int
}

type queue struct {
	data []item
}

func NewQueue() *queue {
	return &queue{
		data: make([]item, 0),
	}
}

func (q *queue) push(d item) {
	q.data = append(q.data, d)
}

func (q *queue) pop() item {
	head := q.data[0]
	q.data = q.data[1:]
	return head
}

func (q *queue) isEmplty() bool {
	return len(q.data) == 0
}

func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	q := NewQueue()
	q.push(item{root, 0})
	for !q.isEmplty() {
		popped := q.pop()
		if len(ans) > popped.level {
			ans[popped.level] = popped.Val
		} else {
			ans = append(ans, popped.Val)
		}

		if popped.Left != nil {
			q.push(item{popped.Left, popped.level + 1})
		}
		if popped.Right != nil {
			q.push(item{popped.Right, popped.level + 1})
		}
	}
	return ans
}
