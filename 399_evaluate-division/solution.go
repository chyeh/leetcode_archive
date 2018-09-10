package solution

func dfs(curr float64, start string, end string, adj map[string]map[string]float64, visited map[string]bool) float64 {
	if _, ok := adj[start]; ok && start == end {
		return curr
	}
	visited[start] = true
	for next, cost := range adj[start] {
		if visited[next] {
			continue
		}
		if ans := dfs(curr*cost, next, end, adj, visited); ans != -1.0 {
			return ans
		}
	}
	visited[start] = false
	return -1.0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adj := make(map[string]map[string]float64, 0)
	for i, equation := range equations {
		if costForward, ok := adj[equation[0]]; ok {
			costForward[equation[1]] = values[i]
		} else {
			costForward = make(map[string]float64, 0)
			costForward[equation[1]] = values[i]
			adj[equation[0]] = costForward
		}

		if costBackward, ok := adj[equation[1]]; ok {
			costBackward[equation[0]] = 1 / values[i]
		} else {
			costBackward = make(map[string]float64, 0)
			costBackward[equation[0]] = 1 / values[i]
			adj[equation[1]] = costBackward
		}
	}
	var ans []float64
	for _, query := range queries {
		visited := make(map[string]bool, 0)
		ans = append(ans, dfs(1.0, query[0], query[1], adj, visited))
	}
	return ans
}
