package solution

// Solution I: Brute Force
/*
func findLength(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	cm := make(map[byte]bool)
	i := 0
	for ; i < len(s); i++ {
		if _, ok := cm[s[i]]; ok {
			return i
		}
		cm[s[i]] = true
	}
	return i
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		l := findLength(s[i:])
		if l > max {
			max = l
		}
	}
	return max
}
*/

// Solution II: Sliding Window
func lengthOfLongestSubstring(s string) int {
	chars := make(map[byte]int)
	l, r := 0, 0
	var max, cnt int
	for r < len(s) {
		if _, ok := chars[s[r]]; !ok {
			chars[s[r]] = 0
		}
		chars[s[r]]++
		if chars[s[r]] > 1 {
			cnt++
		}
		r++

		for cnt > 0 {
			if chars[s[l]] > 1 {
				cnt--
			}
			chars[s[l]]--
			l++
		}

		if r-l > max {
			max = r - l
		}
	}
	return max
}
