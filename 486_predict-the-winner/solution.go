package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PredictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
	}
	for l := 2; l <= len(nums); l++ {
		for i := 0; i < len(nums)-l+1; i++ {
			left := nums[i] - dp[i+1][i+l-1]
			right := nums[i+l-1] - dp[i][i+l-2]
			dp[i][i+l-1] = max(left, right)
		}
	}

	if dp[0][len(nums)-1] >= 0 {
		return true
	}
	return false
}
