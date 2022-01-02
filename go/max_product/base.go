package max_product

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("empty")
	}

	result := nums[0]
	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return result
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("empty")
	}

	result := nums[0]
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

func maxProduct(nums []int) int {
	result, minP, maxP := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxP, minP = max(nums[i], minP*nums[i], maxP*nums[i]), min(nums[i], minP*nums[i], maxP*nums[i])
		result = max(result, maxP)
	}
	return result
}
