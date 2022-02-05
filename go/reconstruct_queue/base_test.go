package reconstruct_queue

import (
	"fmt"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
