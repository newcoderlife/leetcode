package max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	prev, result := 0, 0
	for index, num := range nums {
		if num == 0 {
			prev = index + 1
		} else {
			if index-prev+1 > result {
				result = index - prev + 1
			}
		}
	}

	return result
}
