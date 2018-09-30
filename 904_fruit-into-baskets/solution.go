package solution

func totalFruit(tree []int) int {
	fruits := make(map[int]int)
	l := 0
	var max int
	for r := 0; r < len(tree); r++ {
		if v, ok := fruits[tree[r]]; !ok {
			fruits[tree[r]] = 1
		} else {
			fruits[tree[r]] = v + 1
		}
		for len(fruits) > 2 {
			fruits[tree[l]]--
			if fruits[tree[l]] == 0 {
				delete(fruits, tree[l])
			}
			l++
		}
		if r-l+1 > max {
			max = r - l + 1
		}
	}
	return max
}
