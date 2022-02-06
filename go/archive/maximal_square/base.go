package maximal_square

func min(nums ...int) (result int) {
	result = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < result {
			result = nums[i]
		}
	}
	return
}

func maximalSquare(matrix [][]byte) (result int) {
	buffer := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		buffer = append(buffer, make([]int, len(matrix[0])))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			var left, top, inner int
			if i > 0 {
				top = buffer[i-1][j]
			}
			if j > 0 {
				left = buffer[i][j-1]
			}
			if i > 0 && j > 0 {
				inner = buffer[i-1][j-1]
			}
			buffer[i][j] = min(left, top, inner) + 1
			if buffer[i][j] > result {
				result = buffer[i][j]
			}
		}
	}
	return result * result
}
