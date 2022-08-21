package delete_operation_for_two_strings

func min(nums ...int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if result > nums[i] {
			result = nums[i]
		}
	}
	return result
}

func minDistance(word1 string, word2 string) int {
	if len(word1)*len(word2) == 0 {
		return len(word1) + len(word2)
	}

	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			dp[i][j] = min(dp[i-1][j-1]+2, dp[i-1][j]+1, dp[i][j-1]+1)
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
