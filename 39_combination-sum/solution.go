package solution

func find(candidates []int, target int, sol *[][]int, current []int) {
	if target < 0 {
		return
	}

	cp := make([]int, len(current))
	copy(cp, current)
	if target == 0 {
		*sol = append(*sol, cp)
		return
	}
	for i, n := range candidates {
		find(candidates[i:], target-n, sol, append(cp, n))
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sol := [][]int{}
	for i, n := range candidates {
		current := []int{n}
		find(candidates[i:], target-n, &sol, current)
	}
	return sol
}
