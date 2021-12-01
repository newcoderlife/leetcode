package is_palindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	content := make([]int, 0)
	for current := head; current != nil; current = current.Next {
		content = append(content, current.Val)
	}

	for i := 0; i < len(content); i++ {
		if content[i] != content[len(content)-i-1] {
			return false
		}
	}
	return true
}
