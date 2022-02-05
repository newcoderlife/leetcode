package convert_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	sum = 0
)

func dfs(current *TreeNode) {
	if current == nil {
		return
	}

	dfs(current.Right)

	current.Val += sum
	sum = current.Val

	dfs(current.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}
