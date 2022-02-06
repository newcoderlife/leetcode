package find_duplicate

func findDuplicate(nums []int) int {
	fast, slow := nums[nums[0]], nums[0]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}

	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return fast
}
