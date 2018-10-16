package solution

func longestPalindrome(s string) int {
	chars := make(map[rune]bool)
	var cnt int
	for _, c := range s {
		if _, ok := chars[c]; ok {
			delete(chars, c)
			cnt += 2
		} else {
			chars[c] = true
		}
	}
	if len(chars) > 0 {
		cnt++
	}
	return cnt
}
