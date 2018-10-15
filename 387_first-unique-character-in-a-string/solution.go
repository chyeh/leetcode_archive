package solution

func firstUniqChar(s string) int {
	cnt := make(map[rune]int)
	for _, c := range s {
		if v, ok := cnt[c]; ok {
			cnt[c] = v + 1
		} else {
			cnt[c] = 1
		}
	}
	for i, c := range s {
		if cnt[c] == 1 {
			return i
		}
	}
	return -1
}
