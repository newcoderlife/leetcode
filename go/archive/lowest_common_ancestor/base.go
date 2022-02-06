package lowest_common_ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	result *TreeNode
)

func dfs(current, p, q *TreeNode) bool {
	if current == nil {
		return false
	}

	self := current == p || current == q
	left := dfs(current.Left, p, q)
	right := dfs(current.Right, p, q)

	if (left && right) || (left && self) || (right && self) {
		result = current
		return true
	}
	return left || right || self
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result = nil
	dfs(root, p, q)
	return result
}
