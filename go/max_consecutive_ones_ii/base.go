package max_consecutive_ones_ii

func findMaxConsecutiveOnes(nums []int) int {
	prev, count, result := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		}

		for count > 1 {
			if nums[prev] == 0 {
				count--
			}
			prev++
		}

		if result < i-prev+1 {
			result = i - prev + 1
		}
	}

	return result
}
