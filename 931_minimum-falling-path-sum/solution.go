package solution

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFallingPathSum(A [][]int) int {
	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			p := A[i-1][j]
			if j > 0 {
				p = min(p, A[i-1][j-1])
			}
			if j < len(A[i])-1 {
				p = min(p, A[i-1][j+1])
			}
			A[i][j] = A[i][j] + p
		}
	}
	ans := 10000
	for _, num := range A[len(A)-1] {
		ans = min(ans, num)
	}
	return ans
}
