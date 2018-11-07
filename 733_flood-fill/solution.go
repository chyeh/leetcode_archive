package solution

type ij struct {
	i int
	j int
}

func OOB(a ij, image [][]int) bool {
	if a.i < 0 || a.i >= len(image) {
		return true
	}
	if a.j < 0 || a.j >= len(image[a.i]) {
		return true
	}
	return false
}

type queue struct {
	data []ij
}

func newQueue() *queue {
	return &queue{
		data: make([]ij, 0),
	}
}

func (q *queue) push(d ij) {
	q.data = append(q.data, d)
}

func (q *queue) pop() ij {
	front := q.data[0]
	q.data = q.data[1:]
	return front
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	q := newQueue()
	image[sr][sc] = newColor
	q.push(ij{sr, sc})
	for !q.isEmpty() {
		popped := q.pop()
		directions := []ij{
			{popped.i, popped.j + 1},
			{popped.i, popped.j - 1},
			{popped.i + 1, popped.j},
			{popped.i - 1, popped.j},
		}
		for _, next := range directions {
			if OOB(next, image) || image[next.i][next.j] != oldColor {
				continue
			}
			if image[next.i][next.j] != newColor {
				image[next.i][next.j] = newColor
				q.push(next)
			}
		}
	}
	return image
}
