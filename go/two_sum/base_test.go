package two_sum

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
