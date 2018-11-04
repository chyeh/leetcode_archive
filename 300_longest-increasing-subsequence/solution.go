package solution

// Solution I: Dynamic Programming. Time complexity: O(N^2).
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	var ans int
	dp := make([]int, len(nums))
	for j := 0; j < len(dp); j++ {
		dp[j] = 1
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
		ans = max(ans, dp[j])
	}
	return ans
}

// Solution I: Dynamic Programming with Binary Search. Time complexity: O(NlgN).
// Pending...
