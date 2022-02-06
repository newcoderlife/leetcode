package rob2

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var (
	selectHash map[*TreeNode]int
	passHash   map[*TreeNode]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(current *TreeNode) {
	if current == nil {
		return
	}

	dfs(current.Left)
	dfs(current.Right)

	selectHash[current] = current.Val + passHash[current.Left] + passHash[current.Right]
	passHash[current] = max(selectHash[current.Left], passHash[current.Left]) + max(selectHash[current.Right], passHash[current.Right])
}

func rob(root *TreeNode) int {
	selectHash, passHash = make(map[*TreeNode]int), make(map[*TreeNode]int)
	dfs(root)
	return max(selectHash[root], passHash[root])
}
