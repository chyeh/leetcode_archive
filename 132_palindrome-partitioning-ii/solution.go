package solution

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCut(s string) int {
	p := make([][]bool, len(s))
	for i := range p {
		p[i] = make([]bool, len(s))
	}
	for j := 0; j < len(p); j++ {
		for i := j; i >= 0; i-- {
			if i == j ||
				(j == i+1 && s[i] == s[j]) ||
				(p[i+1][j-1] && s[i] == s[j]) {
				p[i][j] = true
			}
		}
	}

	dp := make([]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = i - 1
	}
	for j := 1; j < len(dp); j++ {
		for i := j - 1; i >= 0; i-- {
			if p[i][j-1] {
				dp[j] = min(dp[j], dp[i]+1)
			}
		}
	}
	return dp[len(s)-1]
}
