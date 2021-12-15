package permute

var (
	results    [][]int
	candidates []int
	visit      []bool
)

func dfs(step int, current []int) {
	if step == len(candidates) {
		result := make([]int, len(current))
		copy(result, current)
		results = append(results, result)
		return
	}
	for i := 0; i < len(candidates); i++ {
		if !visit[i] {
			visit[i] = true
			dfs(step+1, append(current, candidates[i]))
			visit[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	results = [][]int{}
	candidates = nums
	visit = make([]bool, len(candidates))
	dfs(0, []int{})
	return results
}
