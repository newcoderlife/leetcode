package min_path_sum

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	buffer := []int{grid[0][0]}
	for i := 1; i < n; i++ {
		buffer = append(buffer, buffer[i-1]+grid[0][i])
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				buffer[0] = buffer[0] + grid[i][0]
				continue
			}
			buffer[j] = min(buffer[j], buffer[j-1]) + grid[i][j]
		}
	}
	return buffer[n-1]
}
