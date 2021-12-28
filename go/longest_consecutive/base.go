package longest_consecutive

func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for _, num := range nums {
		hash[num] = true
	}

	result := 0
	for _, num := range nums {
		if hash[num-1] {
			continue
		}
		for i := num; ; i++ {
			if !hash[i] {
				break
			}
			if i-num+1 > result {
				result = i - num + 1
			}
		}
	}
	return result
}
