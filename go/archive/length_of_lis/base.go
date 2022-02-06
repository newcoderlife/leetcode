package length_of_lis

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	result := 1
	f := make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		result = max(result, f[i])
	}
	return result
}
