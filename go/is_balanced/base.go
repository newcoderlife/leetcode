package is_balanced

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(current *TreeNode, step int) (bool, int) {
	if current == nil {
		return true, 0
	}
	lResult, lDepth := dfs(current.Left, step+1)
	if !lResult {
		return false, 0
	}
	rResult, rDepth := dfs(current.Right, step+1)
	if !rResult {
		return false, 0
	}
	return abs(lDepth-rDepth) < 2, max(lDepth, rDepth) + 1
}

func isBalanced(root *TreeNode) bool {
	result, _ := dfs(root, 0)
	return result
}
