package max_depth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	return max(dfs(cur.Left), dfs(cur.Right)) + 1
}

func maxDepth(root *TreeNode) int {
	return dfs(root)
}
