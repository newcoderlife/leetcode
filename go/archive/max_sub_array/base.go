package max_sub_array

func maxSubArray(nums []int) int {
	maxNum := nums[0]
	current := 0
	maxSub := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}

		current += num
		if current < 0 {
			current = 0
		} else if current > maxSub {
			maxSub = current
		}
	}

	if maxNum <= 0 {
		return maxNum
	}
	return maxSub
}
