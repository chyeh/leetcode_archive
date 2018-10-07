package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(sum+nums[i], nums[i])
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
