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

func _maxSubArray(nums []int) int {
	result, current := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if current < 0 {
			current = 0
		}
		current += nums[i]
		if current > result {
			result = current
		}
	}

	return result
}
