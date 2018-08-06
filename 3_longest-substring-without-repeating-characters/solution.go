package solution

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
