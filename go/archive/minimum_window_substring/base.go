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

func _minWindow(s string, t string) string {
	target := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}

	result := ""
	count := 0
	current := make(map[byte]int)
	for left, right := 0, 0; right < len(s); right++ {
		current[s[right]]++
		if current[s[right]] == target[s[right]] {
			count++
		}

		for ; left <= right && count == len(target); left++ {
			if len(result) == 0 || right-left+1 < len(result) {
				result = s[left : right+1]
			}
			current[s[left]]--
			if current[s[left]] < target[s[left]] {
				count--
			}
		}
	}

	return result
}
