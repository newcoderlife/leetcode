package k_th_smallest_prime_fraction

type Item struct {
	Value    []int
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

func kthSmallestPrimeFraction(arr []int, k int) []int {
	h := CreateHeap(func(i, j *Item) bool { return i.Priority < j.Priority })
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			h.Push(&Item{Value: []int{arr[i], arr[j]}, Priority: float64(arr[i]) / float64(arr[j])})
		}
	}

	var result []int
	for i := 0; i < k; i++ {
		result = h.Pop().Value
	}
	return result
}

func kthSmallestPrimeFraction2(arr []int, k int) []int {
	h := CreateHeap(func(i, j *Item) bool { return arr[i.Value[0]]*arr[j.Value[1]] < arr[i.Value[1]]*arr[j.Value[0]] })
	for i := 1; i < len(arr); i++ {
		h.Push(&Item{Value: []int{0, i}})
	}

	var result []int
	for i := 0; i < k; i++ {
		result = h.Pop().Value
		if result[0]+1 < result[1] {
			h.Push(&Item{Value: []int{result[0] + 1, result[1]}})
		}
	}

	return []int{arr[result[0]], arr[result[1]]}
}
