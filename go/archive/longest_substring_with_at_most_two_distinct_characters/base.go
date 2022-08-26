package longest_substring_with_at_most_two_distinct_characters

func lengthOfLongestSubstringTwoDistinct(s string) int {
	exist := make(map[byte]int)
	prev, result := 0, 0
	for i := 0; i < len(s); i++ {
		exist[s[i]]++
		if len(exist) > 2 {
			for {
				exist[s[prev]]--
				if exist[s[prev]] == 0 {
					delete(exist, s[prev])
					prev++
					break
				}
				prev++
			}
		}

		if i-prev+1 > result {
			result = i - prev + 1
		}
	}

	return result
}
