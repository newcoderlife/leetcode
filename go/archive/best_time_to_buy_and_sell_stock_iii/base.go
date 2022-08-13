package best_time_to_buy_and_sell_stock_iii

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	// buy1[i] = max(buy1[i-1], -prices[i])
	// sell1[i] = max(sell1[i-1], buy1[i-1]+prices[i])
	// buy2[i] = max(sell1[i-1]-prices[i], buy2[i-1])
	// sell2[i] = max(sell2[i-1], buy2[i-1]+prices[i])
	var (
		buy1         = -prices[0]
		buy2         = -prices[0]
		sell1, sell2 int
	)
	for i := 1; i < len(prices); i++ {
		buy1, sell1, buy2, sell2 = max(buy1, -prices[i]), max(sell1, buy1+prices[i]), max(sell1-prices[i], buy2), max(sell2, buy2+prices[i])
	}
	return sell2
}
