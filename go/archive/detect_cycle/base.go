package detect_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	for fast = head; fast != slow; {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
