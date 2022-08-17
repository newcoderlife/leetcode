package sparse_matrix_multiplication

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m, k, n := len(mat1), len(mat1[0]), len(mat2[0])

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for k1 := 0; k1 < k; k1++ {
			if mat1[i][k1] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				result[i][j] += mat1[i][k1] * mat2[k1][j]
			}
		}
	}

	return result
}
