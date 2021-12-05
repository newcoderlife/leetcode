package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]int, 0), make([]int, 0)
	for current := l1; current != nil; current = current.Next {
		stack1 = append(stack1, current.Val)
	}
	for current := l2; current != nil; current = current.Next {
		stack2 = append(stack2, current.Val)
	}

	head := new(ListNode)
	addon := 0
	for tail1, tail2 := len(stack1)-1, len(stack2)-1; tail1 >= 0 || tail2 >= 0 || addon != 0; tail1, tail2 = tail1-1, tail2-1 {
		head.Next = &ListNode{
			Val:  addon,
			Next: head.Next,
		}
		if tail1 >= 0 {
			head.Next.Val += stack1[tail1]
		}
		if tail2 >= 0 {
			head.Next.Val += stack2[tail2]
		}
		addon = head.Next.Val / 10
		head.Next.Val %= 10
	}
	return head.Next
}
