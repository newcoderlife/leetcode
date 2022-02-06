package num_squares

import (
	"fmt"
	"testing"
)

func Test_numSquares(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Println(i, numSquares(i))
	}
}
