package search_matrix

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if !searchMatrix(matrix, matrix[i][j]) {
				fmt.Println(i, j, matrix[i][j])
			}
		}
	}
}
