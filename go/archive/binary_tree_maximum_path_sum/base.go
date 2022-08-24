package binary_tree_maximum_path_sum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	result int
)

func max(nums ...int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if result < nums[i] {
			result = nums[i]
		}
	}

	return result
}

func dfs(current *TreeNode) int {
	if current == nil {
		return 0
	}

	left, right := dfs(current.Left), dfs(current.Right)

	result = max(result, current.Val, current.Val+left, current.Val+right, current.Val+left+right)

	return max(current.Val, current.Val+left, current.Val+right)
}

func maxPathSum(root *TreeNode) int {
	result = math.MinInt64
	dfs(root)

	return result
}
