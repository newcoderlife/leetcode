package max_area

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left := 0
	right := len(height) - 1
	result := 0
	for left < right {
		result = max(result, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
