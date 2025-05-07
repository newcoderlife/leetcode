package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	for current := head; current != nil; {
		current, current.Next, prev = current.Next, prev, current
	}

	return prev
}
