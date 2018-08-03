package solution

func findSum(nums []int, currSum, S int, ans *int) {
	if len(nums) == 0 || nums == nil {
		return
	}

	if len(nums) == 1 { // last one
		if currSum+nums[0] == S {
			*ans = *ans + 1
		}
		if currSum-nums[0] == S {
			*ans = *ans + 1
		}
		return
	}
	findSum(nums[1:], currSum+nums[0], S, ans)
	findSum(nums[1:], currSum-nums[0], S, ans)
}

func findTargetSumWays(nums []int, S int) int {
	ans := 0
	findSum(nums, 0, S, &ans)
	return ans
}
