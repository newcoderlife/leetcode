package subarray_sum

func subarraySum(nums []int, k int) int {
	var preSum, count int
	hash := map[int]int{0: 1}
	for _, num := range nums {
		preSum += num
		count += hash[preSum-k]
		hash[preSum] += 1
	}
	return count
}
