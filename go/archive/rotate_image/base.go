package rotate_image

func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		for j := 0; j <= (length-1)/2; j++ {
			matrix[i][j], matrix[j][length-i-1], matrix[length-i-1][length-j-1], matrix[length-j-1][i] =
				matrix[length-j-1][i], matrix[i][j], matrix[j][length-i-1], matrix[length-i-1][length-j-1]
		}
	}
}
