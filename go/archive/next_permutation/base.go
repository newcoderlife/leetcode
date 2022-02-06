package next_permutation

import (
	"sort"
)

func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		tIndex := i
		tVal := 1 << 60
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] && nums[j] < tVal {
				tIndex = j
				tVal = nums[j]
			}
		}
		if tIndex != i {
			nums[i], nums[tIndex] = nums[tIndex], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
	}
	sort.Ints(nums)
}
