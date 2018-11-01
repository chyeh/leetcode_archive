package solution

func hasCycle(start, u, v int, adj map[int][]int) bool {
	if v == start {
		return true
	}
	for _, next := range adj[v] {
		if next == u {
			continue
		}
		if hasCycle(start, v, next, adj) {
			return true
		}
	}
	return false
}

func findRedundantConnection(edges [][]int) []int {
	adj := make(map[int][]int)
	for _, e := range edges {
		if hasCycle(e[0], e[0], e[1], adj) {
			return e
		}
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	return nil
}
