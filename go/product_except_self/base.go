package product_except_self

func productExceptSelf(nums []int) []int {
	result := []int{1}
	for i := 1; i < len(nums); i++ {
		result = append(result, nums[i-1]*result[i-1])
	}

	right := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}
