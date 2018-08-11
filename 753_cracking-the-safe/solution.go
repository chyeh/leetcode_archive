package solution

func dfs(n, k int, visited map[string]bool, curr []rune, ans *string, total int) bool {
	if len(visited) == total-1 {
		*ans = string(curr)
		return true
	}
	visited[string(curr[len(curr)-n:])] = true
	for i := 0; i < k; i++ {
		str := string(curr[len(curr)-n+1:]) + string('0'+i)
		if _, ok := visited[str]; !ok {
			curr = append(curr, rune('0'+i))
			if dfs(n, k, visited, curr, ans, total) {
				return true
			}
			curr = curr[:len(curr)-1]
		}
	}
	delete(visited, string(curr[len(curr)-n:]))
	return false
}

func pow(k, n int) int {
	p := 1
	for i := 0; i < n; i++ {
		p *= k
	}
	return p
}

func crackSafe(n int, k int) string {
	powkn := pow(k, n)
	curr := make([]rune, n)
	for i, _ := range curr {
		curr[i] = '0'
	}
	visited := make(map[string]bool)
	var ans string
	for i := 0; i < k; i++ {
		curr[n-1] = rune('0' + i)
		if dfs(n, k, visited, curr, &ans, powkn) {
			break
		}
	}
	return ans
}
