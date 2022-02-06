package num_squares

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	buffer := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prev := 4
		for j := 1; j*j <= i; j++ {
			prev = min(prev, buffer[i-j*j])
		}
		buffer[i] = prev + 1
	}
	return buffer[n]
}
