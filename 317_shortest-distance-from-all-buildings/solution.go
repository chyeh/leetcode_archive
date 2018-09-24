package solution

type coord struct {
	i int
	j int
}

func (c coord) isAccessible(grid [][]int) bool {
	if 0 <= c.i && c.i < len(grid) && 0 <= c.j && c.j < len(grid[0]) && grid[c.i][c.j] == 0 {
		return true
	}
	return false
}

func (c coord) isBuilding(grid [][]int) bool {
	if 0 <= c.i && c.i < len(grid) && 0 <= c.j && c.j < len(grid[0]) && grid[c.i][c.j] == 1 {
		return true
	}
	return false
}

type queue struct {
	data []coord
}

func newQueue() *queue {
	return &queue{
		data: make([]coord, 0),
	}
}

func (q *queue) push(c coord) {
	q.data = append(q.data, c)
}

func (q *queue) pop() coord {
	head := q.data[0]
	q.data = q.data[1:]
	return head
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

func distance(grid [][]int, c coord) [][]int {
	distance := make([][]int, len(grid))
	for i := 0; i < len(distance); i++ {
		distance[i] = make([]int, len(grid[0]))
		for j := 0; j < len(distance[i]); j++ {
			distance[i][j] = -1
		}
	}
	distance[c.i][c.j] = 0

	visited := make([][]bool, len(distance))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(distance[0]))
	}

	q := newQueue()
	visited[c.i][c.j] = true
	q.push(c)

	for !q.isEmpty() {
		popped := q.pop()
		nextSteps := []coord{
			{popped.i, popped.j + 1},
			{popped.i, popped.j - 1},
			{popped.i + 1, popped.j},
			{popped.i - 1, popped.j},
		}
		for _, next := range nextSteps {
			if next.isAccessible(grid) {
				if !visited[next.i][next.j] {
					distance[next.i][next.j] = distance[popped.i][popped.j] + 1
					visited[next.i][next.j] = true
					q.push(next)
				} else if distance[popped.i][popped.j]+1 < distance[next.i][next.j] {
					distance[next.i][next.j] = distance[popped.i][popped.j] + 1
				}
			} else if next.isBuilding(grid) {
				visited[next.i][next.j] = true
				distance[next.i][next.j] = 0
			}
		}
	}
	return distance
}

func shortestDistance(grid [][]int) int {
	buildings := []coord{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				buildings = append(buildings, coord{i, j})
			}
		}
	}

	layers := make([][][]int, 0)
	for _, building := range buildings {
		layer := distance(grid, building)
		for _, b := range buildings {
			if layer[b.i][b.j] == -1 {
				return -1
			}
		}
		layers = append(layers, layer)
	}

	totalDistance := make([][]int, len(grid))
	for i := 0; i < len(totalDistance); i++ {
		totalDistance[i] = make([]int, len(grid[0]))
	}

	for _, layer := range layers {
		for i := 0; i < len(layer); i++ {
			for j := 0; j < len(layer[0]); j++ {
				if totalDistance[i][j] == -1 {
					continue
				}
				if layer[i][j] == -1 {
					totalDistance[i][j] = -1
				} else {
					totalDistance[i][j] += layer[i][j]
				}
			}
		}
	}

	min := 2147483647
	for i := 0; i < len(totalDistance); i++ {
		for j := 0; j < len(totalDistance[0]); j++ {
			if totalDistance[i][j] > 0 && totalDistance[i][j] < min {
				min = totalDistance[i][j]
			}
		}
	}
	if min == 2147483647 {
		return -1
	}
	return min
}
