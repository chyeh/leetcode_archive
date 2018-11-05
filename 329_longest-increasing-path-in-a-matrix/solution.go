package solution

// Time Limit Exceeded
/*
func OOB(matrix [][]int, i, j int) bool {
	if i < 0 || i >= len(matrix) {
		return true
	}
	if j < 0 || j >= len(matrix[0]) {
		return true
	}
	return false
}

func dfs(curr int, matrix [][]int, visited [][]bool, i, j int) int {
	ans := curr
	directions := [][]int{
		{i, j + 1},
		{i, j - 1},
		{i + 1, j},
		{i - 1, j},
	}
	visited[i][j] = true
	for _, next := range directions {
		if OOB(matrix, next[0], next[1]) || visited[next[0]][next[1]] {
			continue
		}
		if matrix[next[0]][next[1]] > matrix[i][j] {
			ans = max(ans, dfs(curr+1, matrix, visited, next[0], next[1]))
		}
	}
	visited[i][j] = false
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestIncreasingPath(matrix [][]int) int {
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	var ans int
	for i := range matrix {
		for j := range matrix[i] {
			ans = max(ans, dfs(1, matrix, visited, i, j))
		}
	}
	return ans
}
*/

// Solution: DFS + Memoization
func OOB(matrix [][]int, i, j int) bool {
	if i < 0 || i >= len(matrix) {
		return true
	}
	if j < 0 || j >= len(matrix[0]) {
		return true
	}
	return false
}

func dfs(matrix [][]int, dp [][]int, i, j int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	dp[i][j] = 1
	directions := [][]int{
		{i, j + 1},
		{i, j - 1},
		{i + 1, j},
		{i - 1, j},
	}
	for _, next := range directions {
		if OOB(matrix, next[0], next[1]) {
			continue
		}
		if matrix[next[0]][next[1]] > matrix[i][j] {
			dp[i][j] = max(dp[i][j], dfs(matrix, dp, next[0], next[1])+1)
		}
	}
	return dp[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	var ans int
	for i := range matrix {
		for j := range matrix[i] {
			ans = max(ans, dfs(matrix, dp, i, j))
		}
	}
	return ans
}
