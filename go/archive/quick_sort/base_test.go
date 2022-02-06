package quick_sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func Test_quickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for {
		total := rand.Intn(1000000)
		buffer := make([]int, total)
		for i := 0; i < total; i++ {
			buffer[i] = int(rand.Int31())
		}

		quickSort(buffer, 0, total)
		if !sort.IntsAreSorted(buffer) {
			t.Errorf("wrong answer")
		}
	}
}
