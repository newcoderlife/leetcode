package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	length := 0
	for current := head; current != nil; current = current.Next {
		length++
	}
	prev, current := (*ListNode)(nil), head
	for count := 0; count*2 < length; count++ {
		current.Next, current, prev = prev, current.Next, current
	}
	if length%2 == 1 {
		prev = prev.Next
	}
	for left, right := prev, current; left != nil && right != nil; left, right = left.Next, right.Next {
		if left.Val != right.Val {
			return false
		}
	}

	return true
}
