package two_sum

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; !ok {
			hash[nums[i]] = i
		}
	}
	for i := 0; i < len(nums); i++ {
		if result, ok := hash[target-nums[i]]; ok && result != i {
			return []int{i, result}
		}
	}
	return nil
}
