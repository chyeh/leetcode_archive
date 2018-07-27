package solution

func solveNQueens(n int) [][]string {
	var sols [][]string
	marks := make([][]bool, n*n)
	for i := 0; i < n; i++ {
		marks[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			marks[i][j] = false
		}
	}
	for j := 0; j < n; j++ {
		find(j, 0, n, marks, []int{j}, &sols)
	}
	return sols
}

func mark(j, row int, n int, marks [][]bool) [][]bool {
	if row >= n {
		return marks
	}

	duplicate := make([][]bool, len(marks))
	for i := range marks {
		duplicate[i] = make([]bool, len(marks[i]))
		copy(duplicate[i], marks[i])
	}

	for r := row; r < n; r++ {
		duplicate[r][j] = true
	}
	if row == n-1 {
		return duplicate
	}
	// if row < n-1
	for r, i := row+1, 1; r < n; r, i = r+1, i+1 {
		if j+i < n {
			duplicate[r][j+i] = true
		}
		if j-i >= 0 {
			duplicate[r][j-i] = true
		}
	}
	return duplicate
}

func find(j int, row int, n int, marks [][]bool, queens []int, sols *[][]string) {
	if row == n-1 {
		*sols = append(*sols, sprintQueens(n, queens))
		return
	}
	duplicateMarks := mark(j, row, n, marks)
	for j := 0; j < n; j++ {
		if !duplicateMarks[row+1][j] {
			find(j, row+1, n, duplicateMarks, append(queens, j), sols)
		}
	}
}

func sprintQueens(n int, queens []int) []string {
	sol := []string{}
	for _, j := range queens {
		b := make([]rune, n)
		for i := range b {
			b[i] = '.'
		}
		b[j] = 'Q'
		s := string(b)
		sol = append(sol, s)
	}
	return sol
}
