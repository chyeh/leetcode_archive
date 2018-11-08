package solution

const intMax = 2147483647

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	distances := make([]int, n)
	for i := range distances {
		distances[i] = intMax
	}
	distances[src] = 0
	for i := 0; i <= K; i++ {
		prev := make([]int, len(distances))
		copy(prev, distances)
		for _, edge := range flights {
			if prev[edge[0]] != intMax && distances[edge[0]]+edge[2] < distances[edge[1]] {
				distances[edge[1]] = prev[edge[0]] + edge[2]
			}
		}
	}

	if distances[dst] == intMax {
		return -1
	}
	return distances[dst]
}
