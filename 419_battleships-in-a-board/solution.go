package solution

func countBattleships(board [][]byte) int {
	var cnt int
	for i, r := range board {
		for j := range r {
			if board[i][j] == '.' {
				continue
			}
			if i > 0 && board[i-1][j] == 'X' {
				continue
			}
			if j > 0 && board[i][j-1] == 'X' {
				continue
			}
			cnt++
		}
	}
	return cnt
}
