package solution

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	ans := make([]int, 0, len(s)-len(p)+1)

	charP := make(map[byte]int)
	for i := range p {
		if v, ok := charP[p[i]]; ok {
			charP[p[i]] = v + 1
		} else {
			charP[p[i]] = 1
		}
	}

	l, r := 0, 0
	cnt := len(p)
	for r < len(s) {
		if v, ok := charP[s[r]]; ok {
			if v > 0 {
				cnt--
			}
			charP[s[r]] = v - 1
		}
		r++

		if cnt == 0 {
			ans = append(ans, l)
		}

		if r-l == len(p) {
			if v, ok := charP[s[l]]; ok {
				if v >= 0 {
					cnt++
				}
				charP[s[l]] = v + 1
			}
			l++
		}
	}
	return ans
}
