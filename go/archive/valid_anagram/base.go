package valid_anagram

func isAnagram(s string, t string) bool {
	exist := make([]int, 26)
	for _, c := range s {
		exist[c-'a']++
	}
	for _, c := range t {
		exist[c-'a']--
	}

	for _, cnt := range exist {
		if cnt != 0 {
			return false
		}
	}
	return true
}
