package largest_rectangle_area

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	result := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			result = max(result, heights[stack[len(stack)-1]]*(i-stack[len(stack)-2]-1))
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for len(stack) > 1 {
		result = max(result, (len(heights)-stack[len(stack)-2]-1)*heights[stack[len(stack)-1]])
		stack = stack[:len(stack)-1]
	}
	return result
}
