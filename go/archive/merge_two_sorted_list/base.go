package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := new(ListNode)
	current := result
	for ; list1 != nil && list2 != nil; current = current.Next {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return result.Next
}
