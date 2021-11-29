package connect

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Record struct {
	node  *Node
	level int
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(&Record{node: root, level: 0})

	for current := queue.Front(); current != nil; current = current.Next() {
		r := current.Value.(*Record)
		if r.node.Left != nil {
			queue.PushBack(&Record{node: r.node.Left, level: r.level + 1})
		}
		if r.node.Right != nil {
			queue.PushBack(&Record{node: r.node.Right, level: r.level + 1})
		}
	}

	for current := queue.Front(); ; current = current.Next() {
		r := current.Value.(*Record)
		if current.Next() == nil {
			return root
		}
		next := current.Next().Value.(*Record)
		if r.level == next.level {
			r.node.Next = next.node
		}
	}
}
