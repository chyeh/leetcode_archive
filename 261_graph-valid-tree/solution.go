package solution

func dfs(n int, visited map[int]bool, adj map[int][]int) {
	visited[n] = true
	for _, next := range adj[n] {
		if !visited[next] {
			visited[next] = true
			dfs(next, visited, adj)
		}
	}
}

func validTree(n int, edges [][]int) bool {
	if n-1 != len(edges) {
		return false
	}
	adj := make(map[int][]int)
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	dfs(0, visited, adj)
	return len(visited) == n
}
