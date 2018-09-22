package solution

func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		dp[i] = make([]string, 0)
		for l := i - 1; l >= 0; l-- {
			lefts := dp[l]
			for _, left := range lefts {
				rights := dp[(i-1)-l]
				for _, right := range rights {
					dp[i] = append(dp[i], "("+left+")"+right)
				}
			}
		}
	}
	return dp[n]
}
