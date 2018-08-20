package solution

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		mrow := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := mrow[board[i][j]]; ok {
				return false
			}
			mrow[board[i][j]] = true
		}
	}

	for j := 0; j < 9; j++ {
		mcol := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := mcol[board[i][j]]; ok {
				return false
			}
			mcol[board[i][j]] = true
		}
	}

	for ii := 0; ii < 9; ii = ii + 3 {
		for jj := 0; jj < 9; jj = jj + 3 {
			msq := make(map[byte]bool)
			for i := ii; i < ii+3; i++ {
				for j := jj; j < jj+3; j++ {
					if board[i][j] == '.' {
						continue
					}
					if _, ok := msq[board[i][j]]; ok {
						return false
					}
					msq[board[i][j]] = true
				}
			}
		}
	}
	return true
}
