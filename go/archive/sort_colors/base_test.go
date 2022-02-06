package sort_colors

import (
	"fmt"
	"testing"
)

func Test_sortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}
