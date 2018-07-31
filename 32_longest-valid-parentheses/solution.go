package solution

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	max := 0
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
			continue
		}

		startIndex := i - dp[i-1] - 1
		if startIndex >= 0 && s[startIndex] == '(' {
			dp[i] = dp[i-1] + 2
		} else {
			dp[i] = 0
			continue
		}

		prevIndex := startIndex - 1
		if prevIndex >= 0 {
			dp[i] += dp[prevIndex]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
