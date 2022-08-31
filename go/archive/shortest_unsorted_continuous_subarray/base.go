package shortest_unsorted_continuous_subarray

import (
	"math"
)

func findUnsortedSubarray(nums []int) int {
	min, max := math.MaxInt64, math.MinInt64
	for i := 1; i < len(nums); i++ {
		if nums[i-1] <= nums[i] {
			continue
		}
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i-1] > max {
			max = nums[i-1]
		}
	}
	if min == math.MaxInt64 && max == math.MinInt64 {
		return 0
	}

	var left, right int
	for left = 0; left < len(nums); left++ {
		if nums[left] > min {
			break
		}
	}
	for right = len(nums) - 1; right >= left; right-- {
		if nums[right] < max {
			break
		}
	}

	return right - left + 1
}
