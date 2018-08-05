package solution

func kInversePairs(n int, k int) int {
	dp := make([][]int, 1001)
	for i := 0; i <= 1000; i++ {
		dp[i] = make([]int, 1001)
	}
	dp[1][0] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			for kk := 0; kk < len(dp[i-1]) && dp[i-1][kk] != 0 && j+kk <= k; kk++ {
				dp[i][j+kk] = (dp[i][j+kk] + dp[i-1][kk]) % 1000000007
			}
		}
	}
	return dp[n][k]
}
