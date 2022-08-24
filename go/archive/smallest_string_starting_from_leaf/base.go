package smallest_string_starting_from_leaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	result string
)

func dfs(node *TreeNode, current string) {
	if node == nil {
		return
	}

	current = string(rune('a'+node.Val)) + current
	if node.Left == nil && node.Right == nil {
		if current < result || result == "" {
			result = current
		}
		return
	}

	dfs(node.Left, current)
	dfs(node.Right, current)
}

func smallestFromLeaf(root *TreeNode) string {
	result = ""
	dfs(root, "")

	return result
}
