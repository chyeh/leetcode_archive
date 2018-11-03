package solution

func numOfElementsSmallerOrEqualTo(matrix [][]int, target int) int {
	var cnt int
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if target >= matrix[i][j] {
			cnt = cnt + (j + 1)
			i++
		} else {
			j--
		}
	}
	return cnt
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	low, high := matrix[0][0], matrix[n-1][n-1]
	for low < high {
		target := low + (high-low)/2
		cnt := numOfElementsSmallerOrEqualTo(matrix, target)
		if cnt < k {
			low = target + 1
		} else {
			high = target
		}
	}
	return low
}
