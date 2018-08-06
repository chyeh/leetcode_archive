package solution

func oob(board [][]byte, pt ij) bool {
	w := len(board)
	l := len(board[0])
	if pt.i < 0 || pt.i >= w || pt.j < 0 || pt.j >= l {
		return true
	}
	return false
}

type ij struct {
	i, j int
}

func dfs(board [][]byte, visited [][]bool, pt ij, word string) bool {
	if board[pt.i][pt.j] != word[0] {
		return false
	}
	if len(word) == 1 {
		return true
	}
	visited[pt.i][pt.j] = true
	directions := []ij{
		{pt.i, pt.j + 1},
		{pt.i, pt.j - 1},
		{pt.i + 1, pt.j},
		{pt.i - 1, pt.j},
	}
	for _, d := range directions {
		if oob(board, d) || visited[d.i][d.j] {
			continue
		}
		if dfs(board, visited, d, word[1:]) {
			return true
		}
	}
	visited[pt.i][pt.j] = false
	return false
}

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(word) == 0 {
		return false
	}
	visited := make([][]bool, len(board))
	for i, _ := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, row := range board {
		for j, _ := range row {
			if board[i][j] == word[0] {
				if dfs(board, visited, ij{i, j}, word) {
					return true
				}
			}
		}
	}
	return false
}
