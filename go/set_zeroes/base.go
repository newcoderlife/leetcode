package set_zeroes

type Record struct {
	x, y int
}

func setZeroes(matrix [][]int) {
	buffer := make([]*Record, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				buffer = append(buffer, &Record{x: i, y: j})
			}
		}
	}

	for _, r := range buffer {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[r.x][i] = 0
		}
		for i := 0; i < len(matrix); i++ {
			matrix[i][r.y] = 0
		}
	}
}
