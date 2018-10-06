package solution

// Solution I: 2 dimensional DP
/*
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for l := 1; l <= len(s); l++ {
		for i := 0; i < len(s)-l+1; i++ {
			if _, ok := words[s[i:i+l]]; ok {
				dp[i][i+l-1] = true
			} else {
				for j := i; j < i+l-1; j++ {
					if dp[i][j] && dp[j+1][i+l-1] {
						dp[i][i+l-1] = true
						break
					}
				}
			}
		}
	}
	return dp[0][len(s)-1]
}
**/

// Solution II: 1 dimensional DP
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}

	dp := make([]bool, len(s)+1)
	for l := 1; l <= len(s); l++ {
		if _, ok := words[s[:l]]; ok {
			dp[l] = true
		} else {
			for i := 1; i < l; i++ {
				if _, ok := words[s[i:l]]; ok && dp[i] {
					dp[l] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}
