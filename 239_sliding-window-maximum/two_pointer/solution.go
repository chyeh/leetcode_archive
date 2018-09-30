package solution

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1)
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if maxIndex <= i-k {
			maxIndex++
			for next := maxIndex + 1; next <= i; next++ {
				if nums[next] > nums[maxIndex] {
					maxIndex = next
				}
			}
		}
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
		if i >= k-1 {
			ans = append(ans, nums[maxIndex])
		}
	}
	return ans
}
