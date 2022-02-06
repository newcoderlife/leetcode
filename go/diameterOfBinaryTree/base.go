package diameterOfBinaryTree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var (
	result int
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func dfs(current *TreeNode) int {
	if current == nil {
		return 0
	}

	left := dfs(current.Left)
	right := dfs(current.Right)
	result = max(result, left+right+1)
	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	result = 0
	dfs(root)
	return result - 1
}
