package best_time_to_buy_and_sell_stock_ii

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	var (
		hold = -prices[0] // 0
		free = 0          // 1
	)

	for i := 1; i < len(prices); i++ {
		// f[i][0] = max(f[i-1][0], f[i-1][1]-prices[i])
		// f[i][1] = max(f[i-1][0]+prices[i], f[i-1][1])
		hold, free = max(hold, free-prices[i]), max(hold+prices[i], free)
	}

	return free
}

func maxProfit2(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans = max(ans, prices[i]-prices[i-1])
	}

	return ans
}
