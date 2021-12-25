package inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	result []int
)

func dfs(current *TreeNode) {
	if current == nil {
		return
	}
	dfs(current.Left)
	result = append(result, current.Val)
	dfs(current.Right)
}

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	dfs(root)
	return result
}
