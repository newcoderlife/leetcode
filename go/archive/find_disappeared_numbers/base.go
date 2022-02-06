package find_disappeared_numbers

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		num := (nums[i] - 1) % n
		nums[num] += n
	}

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= n {
			result = append(result, i+1)
		}
	}
	return result
}
