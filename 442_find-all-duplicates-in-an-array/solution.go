package solution

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findDuplicates(nums []int) []int {
	nums = append([]int{0}, nums...)
	var ans []int
	for _, num := range nums {
		if nums[abs(num)] < 0 {
			ans = append(ans, abs(num))
			continue
		}
		nums[abs(num)] *= -1
	}
	return ans
}
