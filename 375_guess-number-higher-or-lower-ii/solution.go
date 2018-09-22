package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			min := 2147483647
			for c := i; c <= i+l-1; c++ {
				left := dp[i][c-1]
				right := dp[c+1][i+l-1]
				cost := c + max(left, right)
				if cost < min {
					min = cost
				}
			}
			dp[i][i+l-1] = min
		}
	}

	return dp[1][n]
}
