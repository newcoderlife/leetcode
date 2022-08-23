package path_sum_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	exist  map[int]int
	target int
	result int
)

func dfs(current *TreeNode, sum int) {
	if current == nil {
		return
	}

	sum += current.Val
	result += exist[sum-target]
	exist[sum]++

	dfs(current.Left, sum)
	dfs(current.Right, sum)

	exist[sum]--
}

func pathSum(root *TreeNode, targetSum int) int {
	exist = map[int]int{0: 1}
	target = targetSum
	result = 0

	dfs(root, 0)

	return result
}
