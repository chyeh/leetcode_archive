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

func merge(a, b *ListNode) *ListNode {
	ans := &ListNode{}
	dummy := ans
	for a != nil && b != nil {
		if a.Val < b.Val {
			ans.Next = &ListNode{Val: a.Val}
			a = a.Next
		} else {
			ans.Next = &ListNode{Val: b.Val}
			b = b.Next
		}
		ans = ans.Next
	}
	if a == nil {
		ans.Next = b
	} else {
		ans.Next = a
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for _, list := range lists {
		ans = merge(ans, list)
	}
	return ans
}
