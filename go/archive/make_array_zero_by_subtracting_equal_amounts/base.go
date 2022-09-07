package make_array_zero_by_subtracting_equal_amounts

func minimumOperations(nums []int) int {
	counts := make(map[int]bool)
	for _, num := range nums {
		if num != 0 {
			counts[num] = true
		}
	}
	return len(counts)
}
