package solution

func isValidRange(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			if isValidRange(s[l+1:r+1]) || isValidRange(s[l:r]) {
				return true
			}
			return false
		}
	}
	return true
}
