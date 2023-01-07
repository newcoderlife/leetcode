package top_k_frequent_elements

import "container/heap"

type Occur struct {
	Number int
	Times  int
}

type Occurs []Occur

func (o Occurs) Len() int {
	return len(o)
}

func (o Occurs) Less(i, j int) bool {
	return o[i].Times < o[j].Times
}

func (o Occurs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o *Occurs) Push(x interface{}) {
	*o = append(*o, x.(Occur))
}

func (o *Occurs) Pop() interface{} {
	result := (*o)[len(*o)-1]
	*o = (*o)[:len(*o)-1]
	return result
}

func topKFrequentHeap(nums []int, k int) []int {
	if k > len(nums) {
		k = len(nums)
	}

	exist := make(map[int]int)
	for _, num := range nums {
		exist[num]++
	}

	occurs := new(Occurs)
	heap.Init(occurs)
	for num, times := range exist {
		heap.Push(occurs, Occur{Number: num, Times: times})
		if occurs.Len() > k {
			heap.Pop(occurs)
		}
	}

	result := make([]int, 0, k)
	for _, o := range *occurs {
		result = append(result, o.Number)
	}
	return result
}
