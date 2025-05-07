package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	results := make([]int, 0)
	for step := 0; step < (len(matrix)+1)/2 && step < (len(matrix[0])+1)/2; step++ {
		for j := step; j < len(matrix[0])-step; j++ {
			results = append(results, matrix[step][j])
		}
		for i := step + 1; i < len(matrix)-step; i++ {
			results = append(results, matrix[i][len(matrix[0])-step-1])
		}
		if step*2+1 < len(matrix) {
			for j := len(matrix[0]) - step - 2; j >= step; j-- {
				results = append(results, matrix[len(matrix)-step-1][j])
			}
		}
		if step*2+1 < len(matrix[0]) {
			for i := len(matrix) - step - 2; i > step; i-- {
				results = append(results, matrix[i][step])
			}
		}
	}

	return results
}
