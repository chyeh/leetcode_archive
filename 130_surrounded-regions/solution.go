package solution

type coord struct {
	i int
	j int
}

func (c coord) outOfBoard(board [][]byte) bool {
	if 0 <= c.i && c.i < len(board) && 0 <= c.j && c.j < len(board[0]) {
		return false
	}
	return true
}

func dfsIsSurrounded(board [][]byte, c coord) bool {
	board[c.i][c.j] = '.'
	directions := []coord{
		{c.i, c.j + 1},
		{c.i, c.j - 1},
		{c.i + 1, c.j},
		{c.i - 1, c.j},
	}
	for _, next := range directions {
		if next.outOfBoard(board) {
			return false
		}
		if board[next.i][next.j] == 'O' {
			if !dfsIsSurrounded(board, next) {
				return false
			}
		}
	}
	return true
}

func flip(board [][]byte, c coord) {
	board[c.i][c.j] = 'X'
	directions := []coord{
		{c.i, c.j + 1},
		{c.i, c.j - 1},
		{c.i + 1, c.j},
		{c.i - 1, c.j},
	}
	for _, next := range directions {
		if !next.outOfBoard(board) && board[next.i][next.j] == '.' {
			flip(board, next)
		}
	}
}

func recover(board [][]byte, c coord) {
	board[c.i][c.j] = 'O'
	directions := []coord{
		{c.i, c.j + 1},
		{c.i, c.j - 1},
		{c.i + 1, c.j},
		{c.i - 1, c.j},
	}
	for _, next := range directions {
		if !next.outOfBoard(board) && board[next.i][next.j] == '.' {
			recover(board, next)
		}
	}
}

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				if dfsIsSurrounded(board, coord{i, j}) {
					flip(board, coord{i, j})
				} else {
					recover(board, coord{i, j})
				}
			}
		}
	}
}
