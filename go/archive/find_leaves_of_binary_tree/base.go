package find_leaves_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	results [][]int
)

func putResult(count int, value int) {
	for len(results) <= count {
		results = append(results, make([]int, 0))
	}
	results[count] = append(results[count], value)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(current *TreeNode) int {
	if current == nil {
		return 0
	}
	count := max(dfs(current.Left), dfs(current.Right))
	putResult(count, current.Val)
	return count + 1
}

func findLeaves(root *TreeNode) [][]int {
	results = make([][]int, 0)
	dfs(root)
	return results
}
