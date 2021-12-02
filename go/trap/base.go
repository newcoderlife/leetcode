package trap

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax, result := 0, 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			if height[left] < leftMax {
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}
