package permute_unique

import "sort"

var (
	results    [][]int
	candidates []int
	exist      []bool
)

func save(buffer []int) {
	result := make([]int, len(candidates))
	copy(result, buffer)
	results = append(results, result)
}

func dfs(step int, buffer []int) {
	if step == len(candidates) {
		save(buffer)
		return
	}
	for index, candidate := range candidates {
		if exist[index] {
			continue
		}
		if index > 0 && candidates[index] == candidates[index-1] && !exist[index-1] {
			continue
		}

		exist[index] = true
		dfs(step+1, append(buffer, candidate))
		exist[index] = false
	}
}

func permuteUnique(nums []int) [][]int {
	results = make([][]int, 0)
	exist = make([]bool, len(nums))
	candidates = nums
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	dfs(0, []int{})
	return results
}
