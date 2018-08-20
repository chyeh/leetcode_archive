package solution

func findLengthOfLCIS(nums []int) int {
	max := 0
	for i := 0; i < len(nums); {
		j := i + 1
		for ; j < len(nums) && nums[j-1] < nums[j]; j++ {
		}
		if j-i > max {
			max = j - i
		}
		i = j
	}
	return max
}
