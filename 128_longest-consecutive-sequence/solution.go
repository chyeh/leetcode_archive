package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Solution I: Set
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	cnt := 0
	for k := range set {
		delete(set, k)
		lcnt, rcnt := 0, 0
		for l := k - 1; ; l-- {
			if _, ok := set[l]; ok {
				lcnt++
				delete(set, l)
			} else {
				break
			}
		}
		for r := k + 1; ; r++ {
			if _, ok := set[r]; ok {
				rcnt++
				delete(set, r)
			} else {
				break
			}
		}
		cnt = max(cnt, lcnt+rcnt+1)
	}
	return cnt
}

// Solution II: Hash Map
/*
func longestConsecutive(nums []int) int {
	var ans int
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			continue
		}
		var l, r int
		if v, ok := m[num-1]; ok {
			l = v
		}
		if v, ok := m[num+1]; ok {
			r = v
		}
		sum := l + r + 1
		ans = max(ans, sum)
		m[num] = sum
		m[num-l] = sum
		m[num+r] = sum
	}
	return ans
}
*/
