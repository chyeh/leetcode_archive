package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var nodes []*ListNode
	for i := head; i != nil; i = i.Next {
		nodes = append(nodes, i)
	}
	nodes[0].Next = nil
	for i := 1; i < len(nodes); i++ {
		nodes[i].Next = nodes[i-1]
	}
	return nodes[len(nodes)-1]
}
