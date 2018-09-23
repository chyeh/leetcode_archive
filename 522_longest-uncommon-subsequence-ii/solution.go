package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isSubsequence(sub, str string) bool {
	i := 0
	for j := 0; i < len(sub) && j < len(str); j++ {
		if sub[i] == str[j] {
			i++
		}
	}
	return i == len(sub)
}

func findLUSlength(strs []string) int {
	ans := -1
	for i := 0; i < len(strs); i++ {
		j := 0
		for ; j < len(strs); j++ {
			if i == j {
				continue
			}
			if isSubsequence(strs[i], strs[j]) {
				break
			}
		}
		if j == len(strs) {
			ans = max(len(strs[i]), ans)
		}
	}
	return ans
}
