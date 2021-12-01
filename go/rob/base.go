package rob

var (
	memory     []int
	candidates []int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(current int) int {
	if current >= len(candidates) {
		return 0
	}
	if memory[current] != -1 {
		return memory[current]
	}

	memory[current] = max(candidates[current]+dfs(current+2), candidates[current+1]+dfs(current+3))
	return memory[current]
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	candidates = nums
	memory = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memory[i] = -1
	}
	if len(nums) > 0 {
		memory[len(nums)-1] = nums[len(nums)-1]
	}
	if len(nums) > 1 {
		memory[len(nums)-2] = max(nums[len(nums)-1], nums[len(nums)-2])
	}

	return dfs(0)
}
