package solution

// Solution I: Extra space complexity O(N)
/*
func productExceptSelf(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	preArr := make([]int, len(nums))
	postArr := make([]int, len(nums))
	preArr[0], postArr[len(nums)-1] = 1, 1
	for i, j := 1, len(nums)-2; i < len(nums) && j >= 0; i, j = i+1, j-1 {
		preArr[i] = preArr[i-1] * nums[i-1]
		postArr[j] = postArr[j+1] * nums[j+1]
	}
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = preArr[i] * postArr[i]
	}
	return ans
}
*/

// Solution II: Extra space complexity O(1)
func productExceptSelf(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	preArr := make([]int, len(nums))
	preArr[0] = 1
	for i := 1; i < len(nums); i++ {
		preArr[i] = preArr[i-1] * nums[i-1]
	}
	ans := make([]int, len(nums))
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = preArr[i] * right
		right = right * nums[i]
	}
	return ans
}
