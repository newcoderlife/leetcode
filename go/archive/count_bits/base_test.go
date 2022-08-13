package count_bits

import (
	"testing"
)

func Test_countBits(t *testing.T) {
	countBits(2)
	countBits(10000)
	countBits(0)
}
