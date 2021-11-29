package is_valid_bst

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dfs return result, min, max
func dfs(current *TreeNode) (bool, int, int) {
	cMin, cMax := current.Val, current.Val
	if current.Left != nil {
		lResult, lMin, lMax := dfs(current.Left)
		if !lResult || current.Val <= lMax {
			return false, 0, 0
		}
		cMin = min(cMin, lMin)
	}
	if current.Right != nil {
		rResult, rMin, rMax := dfs(current.Right)
		if !rResult || rMin <= current.Val {
			return false, 0, 0
		}
		cMax = max(cMax, rMax)
	}
	return true, cMin, cMax
}

func isValidBST(root *TreeNode) bool {
	result, _, _ := dfs(root)
	return result
}
