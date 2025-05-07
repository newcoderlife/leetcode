package linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast != slow || fast == nil || fast.Next == nil {
		return nil
	}

	for slow = head; fast != slow; {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
