package solution

import "unicode"

// Solution I
/*
func alphanumeric(s string) string {
	for i := 0; i < len(s); {
		if ('a' <= s[i] && s[i] <= 'z') ||
			('A' <= s[i] && s[i] <= 'Z') ||
			('0' <= s[i] && s[i] <= '9') {
			i++
		} else {
			s = s[:i] + s[i+1:]
		}
	}
	return s
}

func isPalindrome(s string) bool {
	s = strings.ToLower(alphanumeric(s))
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
*/

// Solution II
func isAlphanumeric(c byte) bool {
	if ('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		('0' <= c && c <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}
