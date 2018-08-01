package solution

import "container/list"

const MAX_INT int = 2147483647

func oob(w, h int, p ij) bool {
	if p.i >= w || p.i < 0 || p.j >= h || p.j < 0 {
		return true
	}
	return false
}

type ij struct {
	i int
	j int
}

func (p ij) up() ij {
	return ij{p.i + 1, p.j}
}

func (p ij) down() ij {
	return ij{p.i - 1, p.j}
}

func (p ij) left() ij {
	return ij{p.i, p.j - 1}
}

func (p ij) right() ij {
	return ij{p.i, p.j + 1}
}

type queue struct {
	*list.List
}

func newQueue() *queue {
	l := list.New()
	return &queue{l}
}

func (q *queue) push(d interface{}) {
	q.PushBack(d)
}

func (q *queue) pop() interface{} {
	return q.Remove(q.Front())
}

func (q *queue) isEmpty() bool {
	return q.Len() == 0
}

func updateMatrix(matrix [][]int) [][]int {
	w := len(matrix)
	h := len(matrix[0])
	if w == 1 && h == 1 {
		return matrix
	}

	q := newQueue()

	for i, row := range matrix {
		for j, _ := range row {
			if matrix[i][j] == 1 {
				matrix[i][j] = MAX_INT
			} else {
				q.push(ij{i, j})
			}
		}
	}

	for !q.isEmpty() {
		popped := q.pop().(ij)
		directions := []ij{popped.up(), popped.left(), popped.down(), popped.right()}
		for _, d := range directions {
			if !oob(w, h, d) && matrix[d.i][d.j] > matrix[popped.i][popped.j] {
				matrix[d.i][d.j] = matrix[popped.i][popped.j] + 1
				q.push(d)
			}
		}
	}

	return matrix
}
