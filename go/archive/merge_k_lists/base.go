package merge_k_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	heads := make([]*ListNode, 0)
	for _, list := range lists {
		if list != nil {
			heads = append(heads, &ListNode{Next: list})
		}
	}

	result := new(ListNode)
	for current := result; current != nil; current = current.Next {
		minHead := &ListNode{Next: &ListNode{Val: 10001}}
		finish := true
		for _, head := range heads {
			if head.Next != nil && head.Next.Val < minHead.Next.Val {
				minHead = head
				finish = false
			}
		}
		if finish {
			break
		}

		current.Next = minHead.Next
		minHead.Next = minHead.Next.Next
	}
	return result.Next
}
