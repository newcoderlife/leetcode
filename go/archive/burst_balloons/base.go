package burst_balloons

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCoins(nums []int) int {
	values := make([]int, len(nums)+2)
	values[0], values[len(values)-1] = 1, 1
	copy(values[1:], nums)

	dp := make([][]int, len(values))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(dp))
	}

	for i := len(dp) - 2; i >= 0; i-- {
		for j := i + 2; j < len(dp); j++ {
			base := values[i] * values[j]
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+base*values[k])
			}
		}
	}

	return dp[0][len(dp)-1]
}
