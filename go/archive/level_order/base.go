package level_order

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Record struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	results := make([][]int, 0)
	tmp := make([]int, 0)
	buffer := []*Record{{Node: root, Level: 0}}
	for index, level := 0, 0; index < len(buffer); index++ {
		current := buffer[index]
		if current.Level > level {
			level = current.Level
			result := make([]int, len(tmp))
			copy(result, tmp)
			results = append(results, tmp)
			tmp = make([]int, 0)
		}
		tmp = append(tmp, current.Node.Val)
		if current.Node.Left != nil {
			buffer = append(buffer, &Record{Node: current.Node.Left, Level: current.Level + 1})
		}
		if current.Node.Right != nil {
			buffer = append(buffer, &Record{Node: current.Node.Right, Level: current.Level + 1})
		}
	}
	if len(tmp) != 0 {
		results = append(results, tmp)
	}
	return results
}
