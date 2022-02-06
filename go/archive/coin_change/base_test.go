package coin_change

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864))
}
