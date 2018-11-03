package solution

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[0] > nums[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
