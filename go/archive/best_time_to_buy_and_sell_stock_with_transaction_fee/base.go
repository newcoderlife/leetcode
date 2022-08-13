package best_time_to_buy_and_sell_stock_with_transaction_fee

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
	var (
		hold = -prices[0] // 0
		free = 0          // 1
	)

	for i := 1; i < len(prices); i++ {
		// f[i][0] = max(f[i-1][0], f[i-1][1]-prices[i])
		// f[i][1] = max(f[i-1][0]+prices[i], f[i-1][1])
		hold, free = max(hold, free-prices[i]), max(hold+prices[i]-fee, free)
	}

	return free
}

func maxProfit2(prices []int, fee int) int {
	var (
		buy = prices[0] + fee
		ans int
	)

	for i := 1; i < len(prices); i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			ans += prices[i] - buy
			buy = prices[i]
		}
	}

	return ans
}
