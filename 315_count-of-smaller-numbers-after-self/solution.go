package solution

// Solution I: Insertion sort wtih binary search. Time complexity O(NlgN); Space complexity O(N).
func countSmaller(nums []int) []int {
	ans := make([]int, len(nums))
	sorted := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		l, r := 0, len(sorted)
		for l < r {
			m := l + (r-l)/2
			if sorted[m] >= num {
				r = m
			} else {
				l = m + 1
			}
		}
		ans[i] = r
		sorted = append(sorted[:l], append([]int{num}, sorted[l:]...)...)
	}
	return ans
}

// Solution II: Merge sort
// Pending...
