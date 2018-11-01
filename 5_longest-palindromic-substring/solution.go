package solution

// Solution I: 2 dimensional DP. Time Complexity O(n^2). Space Complexity O(n^2).
/*
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	p := make([][]bool, len(s))
	for i := range p {
		p[i] = make([]bool, len(s))
	}
	var maxI, maxJ, maxLen int
	for j := 0; j < len(p); j++ {
		for i := 0; i <= j; i++ {
			if i == j ||
				(j == i+1 && s[i] == s[j]) ||
				p[i+1][j-1] && s[i] == s[j] {
				p[i][j] = true
				if length := (j - i + 1); length > maxLen {
					maxLen = length
					maxI, maxJ = i, j
				}
			}
		}
	}
	return s[maxI : maxJ+1]
}
*/

// Solution II: Expand Around Center
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	longest := 1
	var maxI, maxJ int
	for i := range s {
		for l, r := i, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if currLen := r - l + 1; currLen > longest {
				longest = currLen
				maxI = l
				maxJ = r
			}
		}

		for l, r := i, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if currLen := r - l + 1; currLen > longest {
				longest = currLen
				maxI = l
				maxJ = r
			}
		}
	}
	return s[maxI : maxJ+1]
}
