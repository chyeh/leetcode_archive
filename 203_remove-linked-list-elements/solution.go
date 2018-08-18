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

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	prev, i := head, head.Next
	for ; i != nil; i = i.Next {
		if prev == head && head.Val == val {
			head = head.Next
			prev = prev.Next
			continue
		}
		if i.Val == val {
			prev.Next = i.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}
