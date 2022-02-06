package search_range

import (
	"fmt"
	"sort"
	"testing"
)

func Test_searchRange(t *testing.T) {
	fmt.Println(sort.SearchInts([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}
