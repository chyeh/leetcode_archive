package solution

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp := []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}
