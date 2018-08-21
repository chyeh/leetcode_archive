package solution

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k > len(nums)-1 {
		k = len(nums) - 1
	}
	for i := 1; i < len(nums); i++ {
		for j := max(i-k, 0); j < i; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}
