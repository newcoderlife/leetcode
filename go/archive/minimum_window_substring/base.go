package minimum_window_substring

func minWindow(s string, t string) string {
	target := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}

	cnt := len(target)
	result := s
	window := make(map[byte]int)
	for i, prev := 0, 0; i < len(s); i++ {
		window[s[i]]++
		if window[s[i]] == target[s[i]] {
			cnt--
		}

		for prev < i && target[s[prev]] < window[s[prev]] {
			window[s[prev]]--
			prev++
		}

		if cnt == 0 {
			if len(result) > i-prev+1 {
				result = s[prev : i+1]
			}
		}
	}

	if cnt > 0 {
		return ""
	}
	return result
}
