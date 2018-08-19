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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	middle, tail := head, head.Next
	for tail != nil && tail.Next != nil {
		middle = middle.Next
		tail = tail.Next.Next
	}
	left, right := middle, middle.Next

	prev, iterator := head, head.Next
	for iterator != right {
		nextIterator := iterator.Next
		iterator.Next = prev
		prev = iterator
		iterator = nextIterator
	}
	head.Next = nil

	if tail == nil {
		left = left.Next
	}

	for ; left != nil && right != nil; left, right = left.Next, right.Next {
		if left.Val != right.Val {
			return false
		}
	}
	return true
}
