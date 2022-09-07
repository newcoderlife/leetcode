package smallest_number_in_infinite_set

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	set := Constructor()
	fmt.Println(set.PopSmallest())
	set.AddBack(1)
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
	set.AddBack(2)
	set.AddBack(3)
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
}
