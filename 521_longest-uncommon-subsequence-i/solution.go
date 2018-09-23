package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}
