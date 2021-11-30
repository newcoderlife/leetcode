package postorder_traversal

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var (
	result []int
)

func dfs(current *TreeNode) {
	if current == nil {
		return
	}
	dfs(current.Left)
	dfs(current.Right)
	result = append(result, current.Val)
}

func postorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	dfs(root)
	return result
}
