package codec

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Record struct {
	Node  *TreeNode
	Level int
}

type Codec struct {
}

func Constructor() Codec { return Codec{} }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	level := 0
	queue := list.New()
	queue.PushBack(&Record{Node: root, Level: 0})
	for current := queue.Front(); current != nil; current = current.Next() {
		r := current.Value.(*Record)
		if r.Node == nil {
			continue
		}
		level = r.Level

		queue.PushBack(&Record{Node: r.Node.Left, Level: r.Level + 1})
		queue.PushBack(&Record{Node: r.Node.Right, Level: r.Level + 1})
	}

	builder := strings.Builder{}
	builder.WriteByte('[')
	for current := queue.Front(); current != nil; current = current.Next() {
		r := current.Value.(*Record)
		if r.Level > level {
			builder.WriteByte(']')
			return builder.String()
		}

		if r.Node == nil {
			builder.WriteString("null,")
		} else {
			builder.WriteString(strconv.FormatInt(int64(r.Node.Val), 10))
			builder.WriteByte(',')
		}

		if current == queue.Back() {
			builder.WriteString("]")
		}
	}
	return builder.String()
}

func mustGetNum(content string) int {
	num, _ := strconv.ParseInt(content, 10, 64)
	return int(num)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	commands := strings.Split(data[1:len(data)-1], ",")
	if len(commands) == 0 || commands[0] == "null" || len(commands[0]) == 0 {
		return nil
	}
	if commands[len(commands)-1] == "" {
		commands = commands[:len(commands)-1]
	}

	queue := list.New()
	queue.PushBack(&TreeNode{Val: mustGetNum(commands[0])})
	for index, current := 1, queue.Front(); index < len(commands) && current != nil; current, index = current.Next(), index+2 {
		r := current.Value.(*TreeNode)

		if commands[index] != "null" {
			r.Left = &TreeNode{Val: mustGetNum(commands[index])}
			queue.PushBack(r.Left)
		}
		if index+1 < len(commands) && commands[index+1] != "null" {
			r.Right = &TreeNode{Val: mustGetNum(commands[index+1])}
			queue.PushBack(r.Right)
		}
	}

	return queue.Front().Value.(*TreeNode)
}
