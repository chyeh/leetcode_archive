package solution

func largestSumOfAverages(A []int, K int) float64 {
	p := make([]float64, len(A)+1)
	for i := 0; i < len(A); i++ {
		p[i+1] = p[i] + float64(A[i])
	}

	dp := make([][]float64, len(A))
	for i := range dp {
		dp[i] = make([]float64, K+1)
	}

	for i := 0; i < len(A); i++ {
		dp[i][1] = (p[len(A)] - p[i]) / float64(len(A)-i)
	}
	for k := 2; k <= K; k++ {
		for i := 0; i < len(A); i++ {
			for j := i + 1; j < len(A); j++ {
				sum := ((p[j] - p[i]) / float64(j-i)) + dp[j][k-1]
				if sum > dp[i][k] {
					dp[i][k] = sum
				}
			}
		}
	}

	return dp[0][K]
}
