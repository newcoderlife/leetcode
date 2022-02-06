package climb_stairs

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a, b, f := 1, 2, 0
	for index := 3; index <= n; index++ {
		f = a + b
		a, b = b, f
	}
	return f
}
