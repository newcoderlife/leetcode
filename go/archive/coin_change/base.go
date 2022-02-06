package coin_change

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)
	for index, num := range coins {
		if num > amount {
			coins = coins[:index]
			break
		}
	}

	result := make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		result[coins[i]] = 1
		for j := 1; j+coins[i] <= amount; j++ {
			if result[j] == 0 {
				continue
			}

			if result[j+coins[i]] == 0 {
				result[j+coins[i]] = result[j] + 1
				continue
			}

			result[j+coins[i]] = min(result[j+coins[i]], result[j]+1)
		}
	}
	if result[amount] == 0 {
		return -1
	}
	return result[amount]
}
