package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	results := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		right, target := len(nums)-1, -nums[i]
		for left := i + 1; left < len(nums); left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}

			for left < right && nums[left]+nums[right] > target {
				right--
			}
			if left == right {
				break
			}

			if nums[left]+nums[right] == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
			}
		}
	}

	return results
}
