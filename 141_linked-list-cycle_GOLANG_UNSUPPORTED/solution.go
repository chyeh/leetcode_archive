package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

/* Interpreted to C++ for varification
class Solution {
public:
	bool hasCycle(ListNode *head) {
		if (!head || !head->next) {
			return false;
		}
		ListNode *slow = head, *fast = head->next->next;
		while (fast && fast->next) {
			if (slow == fast) return true;
			slow = slow->next;
			fast = fast->next->next;
		}
		return false;
	}
};
*/
