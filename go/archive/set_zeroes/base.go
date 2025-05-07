package set_zeroes

func setZeroes(matrix [][]int) {
	i0, j0 := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			j0 = true
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			i0 = true
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if i0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if j0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

func _setZeroes(matrix [][]int) {
	line, row := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				line[i] = true
				row[j] = true
			}
		}
	}
	for i := 0; i < len(line); i++ {
		if line[i] {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < len(row); j++ {
		if row[j] {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
}

func __setZeroes(matrix [][]int) {
	line, row := false, false
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			line = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			row = true
			break
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if line {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if row {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
