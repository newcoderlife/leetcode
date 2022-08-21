package one_edit_distance

func isOneEditDistance(s string, t string) bool {
	if len(s) > len(t) {
		return isOneEditDistance(t, s)
	}

	if len(s) != len(t) && len(s)+1 != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			if len(s) == len(t) {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}

	return len(s)+1 == len(t)
}
