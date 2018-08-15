package solution

func numDecodings(s string) int {
	if len(s) == 0 || (len(s) > 0 && s[0] == '0') {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] == '0' {
			dp[i] = 0
		} else if s[i-1] == '*' {
			dp[i] = dp[i-1] * 9
		} else {
			dp[i] = dp[i-1]
		}

		if i > 1 && s[i-1] == '0' {
			if s[i-2] == '*' {
				dp[i] += dp[i-2] * 2
			} else if s[i-2] == '1' || s[i-2] == '2' {
				dp[i] += dp[i-2]
			}
		} else if i > 1 && (s[i-1] == '1' || s[i-1] == '2') {
			if (s[i-2] == '2' && (s[i-1] <= '6')) || s[i-2] == '1' {
				dp[i] += dp[i-2]
			} else if s[i-2] == '*' && s[i-1] <= '6' {
				dp[i] += dp[i-2] * 2
			} else if s[i-2] == '*' && s[i-1] > '6' {
				dp[i] += dp[i-2]
			}
		} else if i > 1 && s[i-1] == '*' {
			if s[i-2] == '2' {
				dp[i] += dp[i-2] * 6
			} else if s[i-2] == '1' {
				dp[i] += dp[i-2] * 9
			} else if s[i-2] == '*' {
				dp[i] += dp[i-2] * 15
			}
		} else if i > 1 {
			if s[i-2] == '2' && s[i-1] <= '6' {
				dp[i] += dp[i-2]
			} else if s[i-2] == '1' {
				dp[i] += dp[i-2]
			} else if s[i-2] == '*' && s[i-1] <= '6' {
				dp[i] += dp[i-2] * 2
			} else if s[i-2] == '*' && s[i-1] > '6' {
				dp[i] += dp[i-2]
			}
		}
		dp[i] %= 1000000007
	}
	return dp[len(s)]
}
