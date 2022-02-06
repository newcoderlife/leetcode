package majority_element

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for index := 1; index < len(nums); index++ {
		if candidate == nums[index] {
			count++
		} else {
			count--
		}
		if count == 0 {
			candidate = nums[index]
			count = 1
		}
	}

	return candidate
}
