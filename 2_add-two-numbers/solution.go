package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	var carry int
	var tail *ListNode
	i, j := l1, l2
	for ; i != nil; i = i.Next {
		a := i.Val
		b := 0
		if j != nil {
			b = j.Val
		}
		sum := a + b + carry
		carry = sum / 10
		i.Val = sum % 10

		if j != nil {
			j = j.Next
		}
		tail = i
	}
	for ; j != nil; j = j.Next {
		sum := j.Val + carry
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return l1
}
