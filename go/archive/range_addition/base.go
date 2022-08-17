package range_addition

func getModifiedArray(length int, updates [][]int) []int {
	result := make([]int, length+1)
	for _, update := range updates {
		result[update[0]] += update[2]
		result[update[1]+1] -= update[2]
	}
	for i := 1; i < length; i++ {
		result[i] += result[i-1]
	}

	return result[:length]
}
