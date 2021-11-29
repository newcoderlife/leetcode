package min_depth

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 {
		if right == 0 {
			return 1
		}
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	return min(left, right) + 1
}
