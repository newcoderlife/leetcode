package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	exist := make(map[byte]int)
	prev, result := 0, 0
	for i := 0; i < len(s); i++ {
		exist[s[i]]++
		if exist[s[i]] == 2 {
			for {
				exist[s[prev]]--
				if exist[s[prev]] == 1 {
					prev++
					break
				}
				delete(exist, s[prev])
				prev++
			}
		}

		if i-prev+1 > result {
			result = i - prev + 1
		}
	}

	return result
}
