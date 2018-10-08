package solution

func dfs(u, v int, adj map[int][]int, visited map[int]bool) bool {
	if u == v {
		return true
	}
	visited[u] = true
	for _, next := range adj[u] {
		if !visited[next] && dfs(next, v, adj, visited) {
			return true
		}
	}
	visited[u] = false
	return false

}

func hasCycle(u, v int, adj map[int][]int) bool {
	visited := make(map[int]bool)
	return dfs(u, v, adj, visited)
}

func findRedundantConnection(edges [][]int) []int {
	adj := make(map[int][]int)
	for _, e := range edges {
		if hasCycle(e[0], e[1], adj) {
			return e
		}
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	return nil
}
