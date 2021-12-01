package is_isomorphic

func isIsomorphic(s string, t string) bool {
	hashAB, hashBA := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		sc, sok := hashAB[s[i]]
		tc, tok := hashBA[t[i]]

		if sok != tok {
			return false
		}
		if sok && tok {
			if sc != t[i] || tc != s[i] {
				return false
			}
		} else {
			hashAB[s[i]] = t[i]
			hashBA[t[i]] = s[i]
		}
	}
	return true
}
