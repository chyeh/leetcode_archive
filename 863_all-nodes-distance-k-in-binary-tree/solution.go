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

type queue struct {
	d []*TreeNode
}

func newQueue() *queue {
	return &queue{
		d: make([]*TreeNode, 0),
	}
}

func (q *queue) push(n *TreeNode) {
	q.d = append(q.d, n)
}

func (q *queue) pop() *TreeNode {
	front := q.d[0]
	q.d = q.d[1:]
	return front
}

func (q *queue) isEmpty() bool {
	return len(q.d) == 0
}

func buildParentMap(root *TreeNode, parent map[int]*TreeNode) {
	if root.Left != nil {
		parent[root.Left.Val] = root
		buildParentMap(root.Left, parent)
	}
	if root.Right != nil {
		parent[root.Right.Val] = root
		buildParentMap(root.Right, parent)
	}
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	parent := make(map[int]*TreeNode)
	buildParentMap(root, parent)

	var ans []int
	distance := make(map[*TreeNode]int)
	visited := make(map[*TreeNode]bool)

	q := newQueue()
	distance[target] = 0
	visited[target] = true
	q.push(target)
	for !q.isEmpty() {
		popped := q.pop()
		if distance[popped] == K {
			ans = append(ans, popped.Val)
			continue
		}
		if popped.Left != nil && !visited[popped.Left] {
			distance[popped.Left] = distance[popped] + 1
			visited[popped.Left] = true
			q.push(popped.Left)
		}
		if popped.Right != nil && !visited[popped.Right] {
			distance[popped.Right] = distance[popped] + 1
			visited[popped.Right] = true
			q.push(popped.Right)
		}
		if p := parent[popped.Val]; p != nil && !visited[p] {
			distance[p] = distance[popped] + 1
			visited[p] = true
			q.push(p)
		}
	}

	return ans
}
