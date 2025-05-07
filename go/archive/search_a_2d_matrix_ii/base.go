package search_a_2d_matrix_ii

func dfs(x, y int, matrix [][]int, target int) bool {
	if matrix[x][y] == target {
		return true
	}
	if x+1 < len(matrix) && target >= matrix[x+1][y] {
		return dfs(x+1, y, matrix, target)
	}
	return y > 0 && dfs(x, y-1, matrix, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	return dfs(0, len(matrix[0])-1, matrix, target)
}

func _searchMatrix(matrix [][]int, target int) bool {
	for x, y := 0, len(matrix[0])-1; x < len(matrix) && y >= 0; {
		if matrix[x][y] == target {
			return true
		}
		if x+1 < len(matrix) && target >= matrix[x+1][y] {
			x++
		} else if y > 0 {
			y--
		} else {
			return false
		}
	}

	return false
}
