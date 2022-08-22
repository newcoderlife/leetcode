package rearrange_string_k_distance_apart

import "container/heap"

type Item struct {
	Value int
	Count int
}
type Heap []*Item

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].Count > h[j].Count }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(*Item)) }
func (h *Heap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}

func rearrangeString(s string, k int) string {
	if k == 0 {
		return s
	}

	counts := make([]int, 26)
	for _, c := range s {
		counts[int(c)-int('a')]++
	}

	h := new(Heap)
	for index, count := range counts {
		if count > 0 {
			heap.Push(h, &Item{Value: index, Count: count})
		}
	}

	buffer := make([]*Item, 0)
	result := make([]byte, 0)
	for h.Len() > 0 {

		item := heap.Pop(h).(*Item)
		result = append(result, byte(item.Value+'a'))
		item.Count--
		buffer = append(buffer, item)
		if len(buffer) == k {
			front := buffer[0]
			buffer = buffer[1:]
			if front.Count > 0 {
				heap.Push(h, front)
			}
		}
	}

	if len(result) == len(s) {
		return string(result)
	}
	return ""
}
