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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	second := head.Next
	oddIterator, evenIterator := head, head.Next
	for oddIterator != nil && oddIterator.Next != nil && oddIterator.Next.Next != nil {
		oddIterator.Next = oddIterator.Next.Next
		oddIterator = oddIterator.Next
		evenIterator.Next = oddIterator.Next
		evenIterator = evenIterator.Next
	}
	oddIterator.Next = second
	return head
}
