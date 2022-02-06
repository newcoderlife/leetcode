package count_bits

func countBits(n int) []int {
	result := []int{0, 1}
	for len(result) <= n {
		current := len(result)
		for i := 0; i < current && len(result) <= n; i++ {
			result = append(result, result[i]+1)
		}
	}
	return result[:n+1]
}
