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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	nodes := []*ListNode{}
	i := 0
	iterator := head
	for ; iterator != nil; i, iterator = i+1, iterator.Next {
		nodes = append(nodes, iterator)
	}
	length := i
	if length < n {
		return head
	}
	node := nodes[length-n]
	if length-n-1 >= 0 {
		prev := nodes[length-n-1]
		prev.Next = node.Next
		return head
	}
	return node.Next
}
