package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	l, i, j := dummy, l1, l2
	for i != nil && j != nil {
		if i.Val < j.Val {
			l.Next = i
			i = i.Next
		} else {
			l.Next = j
			j = j.Next
		}
		l = l.Next
	}
	if i == nil {
		l.Next = j
	} else {
		l.Next = i
	}
	return dummy.Next
}
