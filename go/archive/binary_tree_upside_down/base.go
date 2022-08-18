package binary_tree_upside_down

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(left, top, right *TreeNode) *TreeNode {
	if left == nil {
		return nil
	}

	left1, right1 := left.Left, left.Right
	left.Left, left.Right = right, top

	if left1 == nil {
		return left
	}
	return dfs(left1, left, right1)
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	return dfs(root, nil, nil)
}
