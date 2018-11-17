package solution

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	maxp := nums[0]
	minp := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		maxp, minp =
			max(max(maxp*nums[i], minp*nums[i]), nums[i]),
			min(min(maxp*nums[i], minp*nums[i]), nums[i])
		ans = max(ans, maxp)
	}
	return ans
}
