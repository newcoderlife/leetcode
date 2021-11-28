package combine

var (
	exist      []bool
	candidates []int
	results    [][]int
	length     int
)

func dfs(step int, candidates, buffer []int) {
	if step == length {
		result := make([]int, length)
		copy(result, buffer)
		results = append(results, result)
		return
	}

	for index, num := range candidates {
		if exist[num] {
			continue
		}
		exist[num] = true
		dfs(step+1, candidates[index+1:], append(buffer, num))
		exist[num] = false
	}
}

func combine(n int, k int) [][]int {
	length = k
	exist = make([]bool, n+1)
	results = make([][]int, 0)
	candidates = make([]int, 0)
	for i := 1; i <= n; i++ {
		candidates = append(candidates, i)
	}

	dfs(0, candidates, []int{})
	return results
}
