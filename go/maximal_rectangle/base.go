package maximal_rectangle

func maximalRectangle(matrix [][]byte) int {
	maxLeftLength := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		maxLeftLength = append(maxLeftLength, make([]int, len(matrix[0])))
		if matrix[i][0] == '1' {
			maxLeftLength[i][0] = 1
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				maxLeftLength[i][j] = maxLeftLength[i][j-1] + 1
			}
		}
	}

	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			currentMinLength := 201
			for k := i; k >= 0; k-- {
				if currentMinLength > maxLeftLength[k][j] {
					currentMinLength = maxLeftLength[k][j]
				}
				currentSize := currentMinLength * (i - k + 1)
				if currentSize > result {
					result = currentSize
				}
			}
		}
	}

	return result
}
