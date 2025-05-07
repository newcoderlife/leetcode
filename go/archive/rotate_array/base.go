package rotate_array

func rotate(nums []int, k int) {
	if k %= len(nums); k == 0 {
		return
	}

	copy(nums, append(nums[len(nums)-k:len(nums)], nums[0:len(nums)-k]...))
}

func _rotate(nums []int, k int) {
	if k %= len(nums); k == 0 {
		return
	}

	for start, count := 0, 0; count < len(nums); start++ {
		prev, current := nums[start], start
		for ok := true; ok; ok = current != start {
			next := (current + k) % len(nums)
			nums[next], prev, current = prev, nums[next], next
			count++
		}
	}
}
