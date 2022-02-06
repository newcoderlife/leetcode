package knows

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		result := 0
		for i := 1; i < n; i++ {
			if knows(result, i) {
				result = i
			}
		}
		for i := 0; i < n; i++ {
			if i != result && (!knows(i, result) || knows(result, i)) {
				return -1
			}
		}
		return result
	}
}
