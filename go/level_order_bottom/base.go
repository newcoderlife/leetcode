package level_order_bottom

import "container/list"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Record struct {
	Node  *TreeNode
	Level int
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(&Record{Node: root, Level: 0})

	for current := queue.Front(); current != nil; current = current.Next() {
		r := current.Value.(*Record)
		if r.Node.Left != nil {
			queue.PushBack(&Record{
				Node:  r.Node.Left,
				Level: r.Level + 1,
			})
		}
		if r.Node.Right != nil {
			queue.PushBack(&Record{
				Node:  r.Node.Right,
				Level: r.Level + 1,
			})
		}
	}

	results := make([][]int, 0)
	queue.PushFront(&Record{Node: nil, Level: -1})
	for head, tail := queue.Back(), queue.Back(); head != nil; head = head.Prev() {
		if head.Value.(*Record).Level != tail.Value.(*Record).Level {
			result := make([]int, 0)
			for current := head.Next(); current != tail.Next(); current = current.Next() {
				result = append(result, current.Value.(*Record).Node.Val)
			}
			results = append(results, result)
			tail = head
		}
	}

	return results
}
