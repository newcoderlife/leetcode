package search_range

import "sort"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	lower := sort.SearchInts(nums, target)
	if lower == len(nums) {
		return []int{-1, -1}
	}
	if nums[lower] != target {
		return []int{-1, -1}
	}

	upper := sort.SearchInts(nums, target+1)
	return []int{lower, upper - 1}
}
