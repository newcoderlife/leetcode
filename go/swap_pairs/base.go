package swap_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	realHead := &ListNode{Val: -1, Next: head}
	for current := realHead; current != nil && current.Next != nil; current = current.Next.Next {
		next := current.Next
		next2 := current.Next.Next
		if next2 == nil {
			break
		}
		current.Next = next2
		next.Next = next2.Next
		next2.Next = next
	}
	return realHead.Next
}
