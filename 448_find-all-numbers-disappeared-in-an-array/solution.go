package solution

func findDisappearedNumbers(nums []int) []int {
	nums = append([]int{0}, nums...)
	for i := 1; i < len(nums); i++ {
		for nums[i] != i && nums[nums[i]] != nums[i] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	ans := make([]int, 0, len(nums))
	for i, num := range nums {
		if i != num {
			ans = append(ans, i)
		}
	}
	return ans
}
