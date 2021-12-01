package right_side_view

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Record struct {
	Node  *TreeNode
	Level int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(&Record{Node: root, Level: 0})
	prevLevel :=-1
	results := make([]int, 0)
	for current := queue.Front(); current != nil; current = current.Next() {
		r := current.Value.(*Record)
		if r.Node.Right != nil {
			queue.PushBack(&Record{Node: r.Node.Right, Level: r.Level + 1})
		}
		if r.Node.Left != nil {
			queue.PushBack(&Record{Node: r.Node.Left, Level: r.Level + 1})
		}
		if prevLevel != r.Level {
			results = append(results, r.Node.Val)
			prevLevel = r.Level
		}
	}
	return results
}
