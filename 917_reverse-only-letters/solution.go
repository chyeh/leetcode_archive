package solution

func isLetter(c rune) bool {
	if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
		return false
	}
	return true
}

func reverse(slc []rune) {
	for l, r := 0, len(slc)-1; l < r; l, r = l+1, r-1 {
		slc[l], slc[r] = slc[r], slc[l]
	}
}

func reverseOnlyLetters(S string) string {
	var pure []rune
	for _, c := range S {
		if isLetter(c) {
			pure = append(pure, c)
		}
	}
	reverse(pure)
	var ans []rune
	i := 0
	for _, c := range S {
		if isLetter(c) {
			ans = append(ans, pure[i])
			i++
		} else {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
