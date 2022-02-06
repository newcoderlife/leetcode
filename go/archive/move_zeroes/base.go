package move_zeroes

func moveZeroes(nums []int) {
	prev0 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			prev0++
			continue
		}
		if prev0 != 0 {
			nums[i-prev0] = nums[i]
			nums[i] = 0
		}
	}
}
