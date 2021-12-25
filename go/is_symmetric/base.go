package is_symmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(lhs, rhs *TreeNode) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if lhs == nil || rhs == nil {
		return false
	}

	return lhs.Val == rhs.Val && dfs(lhs.Left, rhs.Right) && dfs(lhs.Right, rhs.Left)
}

func isSymmetric(root *TreeNode) bool {
	return dfs(root, root)
}
