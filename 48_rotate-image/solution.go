package solution

func reverse(row []int) {
	for l, r := 0, len(row)-1; l < r; l, r = l+1, r-1 {
		row[l], row[r] = row[r], row[l]
	}
}

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		reverse(matrix[i])
	}
}
