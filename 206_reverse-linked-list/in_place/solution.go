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
	i, j := head, head.Next
	for j != nil {
		nextJ := j.Next
		j.Next = i
		i = j
		j = nextJ
	}
	head.Next = nil
	return i
}
