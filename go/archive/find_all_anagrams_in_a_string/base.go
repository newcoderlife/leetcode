package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
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

	results := make([]int, 0)
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
			results = append(results, i-len(p)+1)
		}
	}

	return results
}
