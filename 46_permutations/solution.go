package solution

func find(nums []int, curr []int, sols *[][]int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	cp := make([]int, len(curr))
	copy(cp, curr)
	if len(nums) == 1 {
		cp = append(cp, nums...)
		*sols = append(*sols, cp)
	}

	for i, n := range nums {
		nextNums := make([]int, len(nums))
		copy(nextNums, nums)
		nextNums = append(nextNums[:i], nextNums[i+1:]...)
		find(nextNums, append(cp, n), sols)
	}
}

func permute(nums []int) [][]int {
	var sols [][]int
	var curr []int
	find(nums, curr, &sols)
	return sols
}
