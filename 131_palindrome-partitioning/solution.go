package solution

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func findPartition(curr []string, s string, sols *[][]string) {
	if s == "" {
		*sols = append(*sols, curr)
		return
	}
	for i := 1; i <= len(s); i++ {
		cut := s[:i]
		if isPalindrome(cut) {
			next := make([]string, len(curr))
			copy(next, curr)
			findPartition(append(next, cut), s[i:], sols)
		}
	}
}

func partition(s string) [][]string {
	var sols [][]string
	findPartition([]string{}, s, &sols)
	return sols
}
