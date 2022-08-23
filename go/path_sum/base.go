package path_sum

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	buffer  *list.List
	results [][]int
)

func hasPathSum(root *TreeNode, targetSum int, buffer *list.List) {
	if root == nil {
		return
	}

	buffer.PushBack(root.Val)
	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			result := make([]int, 0, buffer.Len())
			for current := buffer.Front(); current != nil; current = current.Next() {
				result = append(result, current.Value.(int))
			}
			results = append(results, result)
		}

		buffer.Remove(buffer.Back())
		return
	}

	hasPathSum(root.Left, targetSum, buffer)
	hasPathSum(root.Right, targetSum, buffer)

	buffer.Remove(buffer.Back())
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	buffer = list.New()
	results = make([][]int, 0)
	hasPathSum(root, targetSum, buffer)
	return results
}
