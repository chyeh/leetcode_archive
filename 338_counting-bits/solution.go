package solution

// Solution I: Most Significant Bit
func countBits(num int) []int {
	dp := make([]int, num+1)
	for b := 1; b <= num; b = b * 2 {
		for i := 0; i < b && b+i <= num; i++ {
			dp[b+i] = dp[i] + 1
		}
	}
	return dp
}

// Solution II: Least Significant Bit
/*
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>2] + (i & 1)
	}
	return dp
}
*/

// Solution III: Last Set Bit
/*
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}
*/
