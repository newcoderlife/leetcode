package permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	return findAnagrams(s2, s1)
}

func findAnagrams(s string, p string) bool {
	window := make([]int, 26)
	for _, c := range p {
		window[c-'a']++
	}

	delta := 0
	for _, count := range window {
		if count > 0 {
			delta++
		}
	}

	for i := 0; i < len(s); i++ {
		window[s[i]-'a']--
		if window[s[i]-'a'] == 0 {
			delta--
		} else if window[s[i]-'a'] == -1 {
			delta++
		}

		if i >= len(p) {
			window[s[i-len(p)]-'a']++
			if window[s[i-len(p)]-'a'] == 0 {
				delta--
			} else if window[s[i-len(p)]-'a'] == 1 {
				delta++
			}
		}

		if delta == 0 {
			return true
		}
	}

	return false
}
