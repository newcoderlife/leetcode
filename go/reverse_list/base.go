package reverse_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, current *ListNode = nil, head
	for current.Next != nil {
		realNext := current.Next
		current.Next = prev
		prev, current = current, realNext
	}
	current.Next = prev
	return current
}
