package sum_of_digits_in_the_minimum_number

import (
	"math"
)

func sumOfDigits(nums []int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	total := 0
	for min > 0 {
		total += min % 10
		min /= 10
	}

	if total%2 == 0 {
		return 1
	}
	return 0
}
