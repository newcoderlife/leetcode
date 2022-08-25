package partition_equal_subset_sum

import "sort"

func canPartition(nums []int) bool {
	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}

	target := sum / 2
	if target*2 != sum || max > target {
		return false
	}

	sort.Ints(nums)
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[target]
}
