package remove_duplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	current := 0
	for _, num := range nums {
		if nums[current] == num {
			continue
		}
		current++
		nums[current] = num
	}
	return current + 1
}
