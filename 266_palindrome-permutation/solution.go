package solution

func canPermutePalindrome(s string) bool {
	chars := make(map[rune]bool)
	for _, c := range s {
		if _, ok := chars[c]; ok {
			delete(chars, c)
		} else {
			chars[c] = true
		}
	}
	if len(chars) < 2 {
		return true
	}
	return false
}
