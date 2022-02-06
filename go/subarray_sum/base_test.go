package subarray_sum

import (
	"fmt"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
