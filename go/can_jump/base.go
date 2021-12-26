package can_jump

func canJump(nums []int) bool {
	current := 0
	for i := 0; i < len(nums); i++ {
		if i > current {
			break
		}

		last := i + nums[i]
		if last >= len(nums) {
			return true
		}
		if last > current {
			current = last
		}
	}
	return current >= len(nums)-1
}
