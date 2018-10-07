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

func maxSubarraySumCircular(A []int) int {
	sum, currMaxSum, currMinSum, maxSum, minSum := A[0], A[0], A[0], A[0], A[0]
	for i := 1; i < len(A); i++ {
		sum += A[i]
		currMaxSum = max(currMaxSum+A[i], A[i])
		maxSum = max(maxSum, currMaxSum)
		currMinSum = min(currMinSum+A[i], A[i])
		minSum = min(minSum, currMinSum)
	}
	if maxSum < 0 {
		return maxSum
	}
	return max(maxSum, sum-minSum)
}
