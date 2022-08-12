package best_time_to_buy_and_sell_stock_with_cooldown

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	var (
		keep   = -prices[0]
		freeze = 0
		free   = 0
	)

	for i := 1; i < len(prices); i++ {
		t_freeze := keep + prices[i]
		keep = max(keep, free-prices[i])
		free = max(freeze, free)
		freeze = t_freeze
	}

	return max(freeze, free)
}
