package solution

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := range dp {
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				smallest := min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
				dp[i][j] = smallest + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
