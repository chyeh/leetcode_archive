package solution

// Recursive: Time Limit Exceeded.
/*
func numDistinct(s string, t string) int {
	if len(t) == 0 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}
	var sum int
	for i := range s {
		if s[i] == t[0] {
			sum += numDistinct(s[i+1:], t[1:])
		}
	}
	return sum
}
*/

// Solution: Dynamic Programming
func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	for j := range dp[0] {
		dp[0][j] = 1
	}
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
