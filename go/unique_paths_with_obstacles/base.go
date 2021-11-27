package unique_paths_with_obstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	buffer := make([]int, n)
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		buffer[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j != 0 {
				buffer[j] += buffer[j-1]
			}
			if obstacleGrid[i][j] == 1 {
				buffer[j] = 0
			}
		}
	}
	return buffer[n-1]
}
