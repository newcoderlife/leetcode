package longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	exist := make([]int, 26)
	prev, max := 0, 0
	for i := 0; i < len(s); i++ {
		exist[s[i]-'A']++
		if exist[s[i]-'A'] > max {
			max = exist[s[i]-'A']
		}

		for i-prev+1 > max+k {
			exist[s[prev]-'A']--
			prev++
		}
	}

	if max+k > len(s) {
		return len(s)
	}
	return max + k
}
