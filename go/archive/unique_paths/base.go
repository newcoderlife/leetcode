package unique_paths

func uniquePaths(m int, n int) int {
	buffer := make([]int, n)
	for i := 0; i < n; i++ {
		buffer[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			buffer[j] += buffer[j-1]
		}
	}
	return buffer[n-1]
}
