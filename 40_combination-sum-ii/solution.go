package solution

import "sort"

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
		if i > 0 && candidates[i-1] == candidates[i] {
			continue
		}
		find(candidates[i+1:], target-n, sol, append(cp, n))
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	sol := [][]int{}
	for i, n := range candidates {
		if i > 0 && candidates[i-1] == candidates[i] {
			continue
		}
		current := []int{n}
		find(candidates[i+1:], target-n, &sol, current)
	}
	return sol
}
