package solution

func isMatch(s string, p string) bool {
	i, j := 0, 0
	istar, jstar := -1, -1
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			istar = i
			jstar = j
			j++
		} else if jstar >= 0 {
			istar++
			i = istar
			j = jstar + 1
		} else {
			return false
		}
	}
	for j < len(p) && p[j] == '*' {
		j++
	}
	if j == len(p) {
		return true
	}
	return false
}
