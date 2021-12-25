package remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	target int
)

func dfs(prev, current *ListNode) int {
	if current == nil {
		return 1
	}

	suffixLen := dfs(current, current.Next)
	if suffixLen == target {
		prev.Next = current.Next
	}
	return suffixLen + 1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	target = n
	pivot := &ListNode{
		Val:  0,
		Next: head,
	}
	dfs(pivot, head)
	return pivot.Next
}
