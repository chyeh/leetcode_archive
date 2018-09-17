package solution

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

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	for n := range rooms {
		visited[n] = false
	}

	q := newQueue()
	visited[0] = true
	q.push(0)
	for !q.isEmpty() {
		popped := q.pop()
		for _, key := range rooms[popped] {
			if !visited[key] {
				visited[key] = true
				q.push(key)
			}
		}
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
