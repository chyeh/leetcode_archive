package solution

func findSols(curr []int, nums []int, sols *[][]int) {
	sol := make([]int, len(curr))
	copy(sol, curr)
	*sols = append(*sols, sol)
	for i := range nums {
		curr = append(curr, nums[i])
		findSols(curr, nums[i+1:], sols)
		curr = curr[:len(curr)-1]
	}
}

func subsets(nums []int) [][]int {
	var sols [][]int
	findSols([]int{}, nums, &sols)
	return sols
}
