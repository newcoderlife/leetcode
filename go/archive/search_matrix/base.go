package search_matrix

func searchMatrix(matrix [][]int, target int) bool {
	valid := func(i, j int) bool {
		return 0 <= i && i < len(matrix) && 0 <= j && j < len(matrix[0])
	}

	for i, j := 0, len(matrix[0])-1; ; {
		for matrix[i][j] < target {
			i++
			if !valid(i, j) {
				return false
			}
		}
		for matrix[i][j] > target {
			j--
			if !valid(i, j) {
				return false
			}
		}
		if matrix[i][j] == target {
			return true
		}
	}
}
