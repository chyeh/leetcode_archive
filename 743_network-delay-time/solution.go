package solution

// BFS doesn't work
/*
type queue struct {
	data []int
}

func newQueue() *queue {
	return &queue{
		data: make([]int, 0),
	}
}

func (q *queue) push(d int) {
	q.data = append(q.data, d)
}

func (q *queue) pop() int {
	front := q.data[0]
	q.data = q.data[1:]
	return front
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

func networkDelayTime(times [][]int, N int, K int) int {
	distances := make(map[int]int)

	q := newQueue()
	distances[K] = 0
	q.push(K)
	for !q.isEmpty() {
		popped := q.pop()
		for _, time := range times {
			if popped == time[0] {
				if v, ok := distances[time[1]]; !ok {
					distances[time[1]] = distances[time[0]] + time[2]
					q.push(time[1])
				} else if distances[time[0]]+time[2] < v {
					distances[time[1]] = distances[time[0]] + time[2]
				}
			}
		}
	}
	var max int
	for i := 1; i <= N; i++ {
		if v, ok := distances[i]; !ok {
			return -1
		} else if v > max {
			max = v
		}
	}
	return max
}
*/

// Solution I: Bellman-Ford Algorithm
/*
const intMax = 2147483647

func networkDelayTime(times [][]int, N int, K int) int {
	distances := make(map[int]int)
	for i := 1; i <= N; i++ {
		distances[i] = intMax
	}
	distances[K] = 0
	for i := 0; i < N; i++ {
		for _, time := range times {
			if distances[time[0]] == intMax {
				continue
			}
			if distances[time[0]]+time[2] < distances[time[1]] {
				distances[time[1]] = distances[time[0]] + time[2]
			}
		}
	}
	var max int
	for _, distance := range distances {
		if distance == intMax {
			return -1
		}
		if distance > max {
			max = distance
		}
	}
	return max
}
*/

// Solution II: Dijkstra's Algorithm
const intMax = 2147483647

func networkDelayTime(times [][]int, N int, K int) int {
	graph := make(map[int]bool)
	for i := 1; i <= N; i++ {
		graph[i] = true
	}

	distances := make(map[int]int)
	for i := 1; i <= N; i++ {
		distances[i] = intMax
	}

	outDegreeEdges := make(map[int][]edge)
	for _, time := range times {
		if _, ok := outDegreeEdges[time[0]]; !ok {
			outDegreeEdges[time[0]] = []edge{edge{v: time[1], weight: time[2]}}
		} else {
			outDegreeEdges[time[0]] = append(outDegreeEdges[time[0]], edge{v: time[1], weight: time[2]})
		}
	}

	V := make([]vertex, 0)
	pq := newPriorityQueue(V)

	distances[K] = 0
	pq.insert(vertex{number: K, distance: 0})
	for len(graph) != 0 && !pq.isEmpty() {
		popped := pq.popMin()
		if _, ok := graph[popped.number]; !ok {
			continue
		}
		delete(graph, popped.number)
		for _, edge := range outDegreeEdges[popped.number] {
			if popped.distance+edge.weight < distances[edge.v] {
				distances[edge.v] = popped.distance + edge.weight
			}
			pq.insert(vertex{number: edge.v, distance: popped.distance + edge.weight})
		}
	}
	var max int
	for _, distance := range distances {
		if distance == intMax {
			return -1
		}
		if distance > max {
			max = distance
		}
	}
	return max
}

type edge struct {
	v      int
	weight int
}

type vertex struct {
	number   int
	distance int
}

func (v vertex) closerThan(a vertex) bool {
	return v.distance < a.distance
}

type priorityQueue struct {
	V []vertex
}

func newPriorityQueue(V []vertex) *priorityQueue {
	pq := &priorityQueue{
		V: V,
	}
	pq.buildMinHeap()
	return pq
}

func (pq *priorityQueue) buildMinHeap() {
	for i := (len(pq.V) - 2) / 2; i >= 0; i-- {
		pq.heapify(i)
	}
}

func (pq *priorityQueue) heapify(i int) {
	smallest := i
	l, r := (2*i)+1, (2*i)+2
	if l < len(pq.V) && pq.V[l].closerThan(pq.V[smallest]) {
		smallest = l
	}
	if r < len(pq.V) && pq.V[r].closerThan(pq.V[smallest]) {
		smallest = r
	}
	if smallest != i {
		pq.V[i], pq.V[smallest] = pq.V[smallest], pq.V[i]
		pq.heapify(smallest)
	}
}

func (pq *priorityQueue) insert(v vertex) {
	pq.V = append(pq.V, v)
	for p, i := (len(pq.V)-2)/2, len(pq.V)-1; i > 0 && pq.V[i].closerThan(pq.V[p]); p, i = (p-1)/2, (i-1)/2 {
		pq.V[i], pq.V[p] = pq.V[p], pq.V[i]
	}
}

func (pq *priorityQueue) popMin() vertex {
	min := pq.V[0]
	pq.V[0], pq.V[len(pq.V)-1] = pq.V[len(pq.V)-1], pq.V[0]
	pq.V = pq.V[:len(pq.V)-1]
	pq.heapify(0)
	return min
}

func (pq *priorityQueue) isEmpty() bool {
	return len(pq.V) == 0
}
