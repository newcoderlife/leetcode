package closest_binary_search_tree_value_ii

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Value    int
	Priority float64
}

// Heap implements small heap
type Heap struct {
	buffer []*Item
	len    int
	less   func(i, j *Item) bool
}

func CreateHeap(less func(i, j *Item) bool) *Heap {
	return &Heap{
		buffer: make([]*Item, 0),
		less:   less,
	}
}

func (h *Heap) Push(item *Item) {
	if len(h.buffer) == h.len {
		h.buffer = append(h.buffer, item)
	} else {
		h.buffer[h.len] = item
	}
	h.up(h.len)
	h.len++
}

func (h *Heap) Top() *Item {
	return h.buffer[0]
}

func (h *Heap) Pop() *Item {
	result := h.buffer[0]
	h.buffer[h.len-1], h.buffer[0] = h.buffer[0], h.buffer[h.len-1]
	h.len--
	h.down(0)
	return result
}

func (h Heap) Len() int {
	return h.len
}

func (h *Heap) down(index int) bool {
	current := index
	for {
		next := 2*current + 1
		if next >= h.len || next < 0 {
			break
		}
		if next2 := next + 1; next2 < h.len && h.less(h.buffer[next2], h.buffer[next]) {
			next = next2
		}

		if h.less(h.buffer[current], h.buffer[next]) {
			break
		}
		h.buffer[current], h.buffer[next] = h.buffer[next], h.buffer[current]
		current = next
	}
	return current > index
}

func (h *Heap) up(current int) {
	for {
		next := (current - 1) / 2
		if next == current || h.less(h.buffer[next], h.buffer[current]) {
			break
		}
		h.buffer[next], h.buffer[current] = h.buffer[current], h.buffer[next]
		current = next
	}
}

func dfs(current *TreeNode, h *Heap, target float64, k int) {
	if current == nil {
		return
	}

	h.Push(&Item{Value: current.Val, Priority: math.Abs(float64(current.Val) - target)})
	dfs(current.Left, h, target, k)
	dfs(current.Right, h, target, k)
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	h := CreateHeap(func(i, j *Item) bool { return i.Priority < j.Priority })
	dfs(root, h, target, k)
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, h.Pop().Value)
	}
	return result
}
