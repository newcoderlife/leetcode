package max_profit

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	var (
		prevMin = 1 << 60
		result  = 0
	)
	for _, num := range prices {
		result = max(result, num-prevMin)
		prevMin = min(prevMin, num)
	}
	return result
}
