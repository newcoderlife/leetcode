package count_substrings_with_only_one_distinct_letter

func countLetters(s string) int {
	result := 0
	for start, end := 0, 0; start < len(s); start = end {
		for end = start + 1; end < len(s) && s[end] == s[start]; end++ {
		}

		for i := 1; i <= end-start; i++ {
			result += i
		}
	}

	return result
}
