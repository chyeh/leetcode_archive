package solution

func searchInsert(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
		if target < nums[i] {
			break
		}
	}
	return i
}
