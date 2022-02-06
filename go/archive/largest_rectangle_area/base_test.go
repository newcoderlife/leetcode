package largest_rectangle_area

import (
	"fmt"
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{4, 2, 0, 3, 2, 5})) // 6
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))          // 3
}
