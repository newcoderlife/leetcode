package sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	result int
)

func dfs(node *TreeNode, current int) {
	if node == nil {
		return
	}

	current = current*10 + node.Val
	if node.Left == nil && node.Right == nil {
		result += current
	}

	dfs(node.Left, current)
	dfs(node.Right, current)
}

func sumNumbers(root *TreeNode) int {
	result = 0
	dfs(root, 0)

	return result
}
