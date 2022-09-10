package the_k_weakest_rows_in_a_matrix

type Item struct {
	Value    int
	Priority int
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

func (h *Heap) Pop() *Item {
	result := h.buffer[0]
	h.buffer[h.len-1], h.buffer[0] = h.buffer[0], h.buffer[h.len-1]
	h.len--
	h.down(0)
	return result
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

func kWeakestRows(mat [][]int, k int) []int {
	h := CreateHeap(func(i, j *Item) bool {
		if i.Priority < j.Priority {
			return true
		}
		if i.Priority == j.Priority && i.Value < j.Value {
			return true
		}
		return false
	})
	for i := 0; i < len(mat); i++ {
		count := 0
		for j := 0; j < len(mat[i]) && mat[i][j] == 1; j++ {
			count++
		}
		h.Push(&Item{Value: i, Priority: count})
	}

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, h.Pop().Value)
	}
	return result
}
