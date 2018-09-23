package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	tortoise, hare := head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
		if tortoise == hare {
			break
		}
	}
	if tortoise != hare {
		return nil
	}
	i, j := head, hare
	for i != j {
		i = i.Next
		j = j.Next
	}
	return i
}

/* Interpreted to C++ for varification
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        if (!head || !head->next) {
            return NULL;
        }
        ListNode *tortoise = head, *hare = head;
        while (hare && hare->next) {
            tortoise = tortoise->next;
            hare = hare->next->next;
            if (tortoise == hare) break;
        }
        if (tortoise != hare) return NULL;
        ListNode *i = head, *j = hare;
        while (i != j) {
            i = i->next;
            j = j->next;
        }
        return i;
    }
};
*/
