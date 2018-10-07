package solution

func isMatched(a, b string) bool {
	charsA := make(map[byte]byte)
	charsB := make(map[byte]byte)
	for i := 0; i < len(a); i++ {
		ca, aok := charsA[a[i]]
		cb, bok := charsB[b[i]]
		if !aok && !bok {
			charsA[a[i]] = b[i]
			charsB[b[i]] = a[i]
		} else if aok && bok {
			if ca != b[i] || cb != a[i] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	ans := make([]string, 0, len(words))
	for _, w := range words {
		if isMatched(w, pattern) {
			ans = append(ans, w)
		}
	}
	return ans
}
