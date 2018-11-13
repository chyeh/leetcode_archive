package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Solution I: 2 dimensional DP
func maximalSquare(matrix [][]byte) int {
	var maxLen int
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
				}
			} else if matrix[i][j] == '1' {
				l := dp[i-1][j-1]
				l = min(l, dp[i][j-1])
				l = min(l, dp[i-1][j])
				dp[i][j] = l + 1
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}
	return maxLen * maxLen
}

// Solution II: 1 dimensional DP
// Pending...
