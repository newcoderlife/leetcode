package merge_two_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := new(ListNode)
	for current, current1, current2 := list, list1, list2; current1 != nil || current2 != nil; {
		if current1 == nil {
			current.Next = &ListNode{Val: current2.Val}
			current2 = current2.Next
		} else if current2 == nil {
			current.Next = &ListNode{Val: current1.Val}
			current1 = current1.Next
		} else if current1.Val > current2.Val {
			current.Next = &ListNode{Val: current2.Val}
			current2 = current2.Next
		} else {
			current.Next = &ListNode{Val: current1.Val}
			current1 = current1.Next
		}
		current = current.Next
	}
	return list.Next
}
