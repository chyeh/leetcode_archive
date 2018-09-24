package solution

type queue struct {
	data []coord
}

func newQueue() *queue {
	return &queue{
		data: make([]coord, 0),
	}
}

func (q *queue) push(d coord) {
	q.data = append(q.data, d)
}

func (q *queue) pop() coord {
	head := q.data[0]
	q.data = q.data[1:]
	return head
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

type coord struct {
	i int
	j int
}

func (c coord) isAccessible(rooms [][]int) bool {
	if 0 <= c.i && c.i < len(rooms) && 0 <= c.j && c.j < len(rooms[0]) && rooms[c.i][c.j] > 0 {
		return true
	}
	return false
}

func wallsAndGates(rooms [][]int) {
	q := newQueue()
	for i, row := range rooms {
		for j := range row {
			if rooms[i][j] == 0 {
				q.push(coord{i, j})
			}
		}
	}

	for !q.isEmpty() {
		popped := q.pop()
		directions := []coord{
			{popped.i, popped.j + 1},
			{popped.i, popped.j - 1},
			{popped.i + 1, popped.j},
			{popped.i - 1, popped.j},
		}
		for _, next := range directions {
			if next.isAccessible(rooms) && rooms[next.i][next.j] > rooms[popped.i][popped.j]+1 {
				rooms[next.i][next.j] = rooms[popped.i][popped.j] + 1
				q.push(next)
			}
		}
	}
}
