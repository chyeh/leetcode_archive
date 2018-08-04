package solution

import (
	"container/list"
)

func oob(grid [][]byte, pt ij) bool {
	w := len(grid)
	l := len(grid[0])
	if pt.i < 0 || pt.i >= w || pt.j < 0 || pt.j >= l {
		return true
	}
	return false
}

type ij struct {
	i int
	j int
}

type queue struct {
	*list.List
}

func NewQueue() *queue {
	return &queue{
		list.New(),
	}
}

func (q *queue) push(d interface{}) {
	q.List.PushBack(d)
}

func (q *queue) pop() interface{} {
	return q.Remove(q.List.Front())
}

func (q *queue) isEmpty() bool {
	return q.Len() == 0
}

func bfs(grid [][]byte, visited [][]bool, i, j int) {
	q := NewQueue()
	visited[i][j] = true
	q.push(ij{i, j})

	for !q.isEmpty() {
		popped := q.pop().(ij)
		directions := []ij{
			{popped.i + 1, popped.j},
			{popped.i - 1, popped.j},
			{popped.i, popped.j + 1},
			{popped.i, popped.j - 1},
		}
		for _, d := range directions {
			if !oob(grid, d) && grid[d.i][d.j] == '1' && !visited[d.i][d.j] {
				visited[d.i][d.j] = true
				q.push(d)
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	w := len(grid)
	l := len(grid[0])

	visited := make([][]bool, w)
	for i, _ := range visited {
		visited[i] = make([]bool, l)
	}

	for i, row := range visited {
		for j, _ := range row {
			visited[i][j] = false
		}
	}

	var num int
	for i, row := range grid {
		for j, _ := range row {
			if visited[i][j] {
				continue
			}

			if grid[i][j] == '1' {
				num++
				bfs(grid, visited, i, j) // find all adjacent 1 and marks visited.
			}
		}
	}
	return num
}
