package solution

func find(sol []int, candidates []int, target int, sols *[][]int) {
	if target < 0 {
		return
	}
	cp := make([]int, len(sol))
	copy(cp, sol)
	if target == 0 {
		*sols = append(*sols, cp)
		return
	}
	for i, candidate := range candidates {
		find(append(cp, candidate), candidates[i:], target-candidate, sols)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var sols [][]int
	find([]int{}, candidates, target, &sols)
	return sols
}
