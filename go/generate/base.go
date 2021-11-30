package generate

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	results := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		result := make([]int, i+1)
		result[0] = 1
		result[i] = 1
		for j := 1; j < i; j++ {
			result[j] = results[i-1][j-1] + results[i-1][j]
		}
		results = append(results, result)
	}
	return results
}
