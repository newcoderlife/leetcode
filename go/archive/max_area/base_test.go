package max_area

import (
	"fmt"
	"testing"
)

func Test_maxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 2, 1}))
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}
