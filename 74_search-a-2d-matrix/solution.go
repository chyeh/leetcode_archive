package solution

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rhead, rtail := 0, len(matrix)-1
	for rhead <= rtail {
		mRow := (rhead + rtail) / 2
		if matrix[mRow][0] <= target && target <= matrix[mRow][len(matrix[mRow])-1] {
			l, r := 0, len(matrix[mRow])-1
			for l <= r {
				m := (l + r) / 2
				if target == matrix[mRow][m] {
					return true
				} else if target > matrix[mRow][m] {
					l = m + 1
				} else {
					r = m - 1
				}
			}
			return false
		} else if target > matrix[mRow][len(matrix[mRow])-1] {
			rhead = mRow + 1
		} else {
			rtail = mRow - 1
		}
	}
	return false
}
