package solution

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charS := make(map[byte]byte)
	charT := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		sc, sok := charS[s[i]]
		tc, tok := charT[t[i]]
		if (sok && !tok) || (!sok && tok) {
			return false
		}
		if sok && tok {
			if sc != t[i] || tc != s[i] {
				return false
			}
		} else {
			charS[s[i]] = t[i]
			charT[t[i]] = s[i]
		}
	}
	return true
}
