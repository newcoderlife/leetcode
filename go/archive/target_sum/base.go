package target_sum

var (
	numbers []int
	memory  map[int]int
)

func dfs(index, target int) int {
	key := (index << 20) + target
	if index == len(numbers) {
		if target == 0 {
			memory[key] = 1
			return 1
		}
		memory[key] = 0
		return 0
	}

	if result, ok := memory[key]; ok {
		return result
	}

	result := dfs(index+1, target-numbers[index]) + dfs(index+1, target+numbers[index])
	memory[key] = result
	return result
}

func findTargetSumWays(nums []int, target int) int {
	numbers = nums
	memory = make(map[int]int)
	return dfs(0, target)
}
