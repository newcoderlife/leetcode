package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(lhs, rhs *ListNode) *ListNode {
	pivot := new(ListNode)
	lp, rp, cur := lhs, rhs, pivot
	for ; lp != nil && rp != nil; cur = cur.Next {
		if lp.Val < rp.Val {
			cur.Next = lp
			lp = lp.Next
		} else {
			cur.Next = rp
			rp = rp.Next
		}
	}
	if lp != nil {
		cur.Next = lp
	} else {
		cur.Next = rp
	}
	return pivot.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	fast, slow := head, head
	for fast != tail {
		fast, slow = fast.Next, slow.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	return merge(sort(head, slow), sort(slow, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}
