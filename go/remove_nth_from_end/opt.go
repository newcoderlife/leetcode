package remove_nth_from_end

func removeNthFromEndOpt(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	pivot := &ListNode{
		Val:  0,
		Next: head,
	}
	fast, slow := pivot, pivot
	for i := 0; i <= n && fast != nil; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return pivot.Next
}
