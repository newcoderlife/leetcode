package combination_sum2

import (
	"sort"
)

var (
	results, nums [][]int
)

func dfs(step, target int, buffer []int) {
	if target == 0 {
		result := make([]int, len(buffer))
		copy(result, buffer)
		results = append(results, result)
		return
	}
	if step == len(nums) || target < 0 {
		return
	}

	current := len(buffer)
	for take := 0; take <= nums[step][1]; take++ {
		for i := 0; i < take; i++ {
			buffer = append(buffer, nums[step][0])
		}
		dfs(step+1, target-nums[step][0]*take, buffer)
		buffer = buffer[:current]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	results, nums = make([][]int, 0), make([][]int, 0)
	count := make(map[int]int)
	for _, candidate := range candidates {
		count[candidate] += 1
	}
	for key, value := range count {
		nums = append(nums, []int{key, value})
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	dfs(0, target, []int{})
	return results
}
