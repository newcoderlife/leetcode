package best_time_to_buy_and_sell_stock_iv

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	k = min(k, n/2)
	sell := make([][]int, n)
	buy := make([][]int, n)
	for i := 0; i < n; i++ {
		sell[i] = make([]int, k+1)
		buy[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt32
		sell[0][i] = math.MinInt32
	}

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}

	ans := 0
	for i := 0; i <= k; i++ {
		ans = max(ans, sell[n-1][i])
	}
	return ans
}
