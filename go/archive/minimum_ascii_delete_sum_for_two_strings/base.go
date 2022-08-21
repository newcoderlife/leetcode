package minimum_ascii_delete_sum_for_two_strings

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeleteSum(word1 string, word2 string) int {
	if len(word1)*len(word2) == 0 {
		return len(word1) + len(word2)
	}

	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = dp[i-1][0] + int(word1[i-1])
	}
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = dp[0][j-1] + int(word2[j-1])
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(word1[i-1]), dp[i][j-1]+int(word2[j-1]))
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
