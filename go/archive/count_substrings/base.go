package count_substrings

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		for l, r := i, i; 0 <= l && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			count++
		}
		for l, r := i-1, i; 0 <= l && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			count++
		}
	}
	return count
}
