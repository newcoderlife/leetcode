package the_k_weakest_rows_in_a_matrix

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestCreateHeap(t *testing.T) {
	n := int(1e7)
	h := CreateHeap(func(i, j *Item) bool { return i.Priority < j.Priority })
	buffer := make([]int, 0)
	for i := 0; i < n; i++ {
		num := rand.Int()
		buffer = append(buffer, num)
		h.Push(&Item{Priority: num})
	}
	sort.Ints(buffer)

	count := 0
	for i := 0; i < n; i++ {
		if h.Pop().Priority != buffer[i] {
			count++
		}
	}
	fmt.Println(count)
}
