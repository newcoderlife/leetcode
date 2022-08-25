package partition_to_k_equal_sum_subsets

import "sort"

var (
	memory map[int]bool
)

func dfs(nums []int, current int, bucket []int, k int, target int, exist int) bool {
	if k == 0 {
		return true
	}

	if bucket[k] == target {
		memory[exist] = dfs(nums, 0, bucket, k-1, target, exist)
		return memory[exist]
	}

	if result, ok := memory[exist]; ok {
		return result
	}

	for i := current; i < len(nums); i++ {
		if exist&(1<<i) > 0 {
			continue
		}
		if bucket[k]+nums[i] > target {
			memory[exist] = false
			return false
		}

		exist |= 1 << i
		bucket[k] += nums[i]
		result := dfs(nums, current+1, bucket, k, target, exist)
		bucket[k] -= nums[i]
		exist ^= 1 << i

		if result {
			memory[exist] = true
			return true
		} else {
			for i+1 < len(nums) && nums[i+1] == nums[i] {
				i++
			}
		}
	}

	memory[exist] = false
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}

	target := sum / k
	if target*k != sum || target < max {
		return false
	}

	sort.Ints(nums)
	bucket := make([]int, k+1)
	memory = make(map[int]bool)
	return dfs(nums, 0, bucket, k, target, 0)
}

func canPartitionKSubsets2(nums []int, k int) bool {
	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}

	target := sum / k
	if target*k != sum || target < max {
		return false
	}

	n := len(nums)
	buffer := make([]int, 1<<n)
	for i := 1; i < len(buffer); i++ {
		buffer[i] = -1
	}

	for state := 1; state < (1 << n); state++ {
		for index, num := range nums {
			if state&(1<<index) == 0 {
				continue
			}

			prevState := state & ^(1 << index)
			if buffer[prevState] >= 0 && buffer[prevState]+num <= target {
				buffer[state] = (buffer[prevState] + num) % target
			}
		}
	}

	return buffer[1<<n-1] == 0
}
