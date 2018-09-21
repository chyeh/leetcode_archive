package solution

func new21Game(N int, K int, W int) float64 {
	dp := make([]float64, N+1)
	dp[0] = 1.0
	var sum float64
	for i := 1; i <= N; i++ {
		if i <= K {
			sum += dp[i-1]
		}
		if i > W {
			sum -= dp[i-W-1]
		}
		dp[i] += (sum / float64(W))
	}
	var p float64
	for i := K; i <= N; i++ {
		p += dp[i]
	}
	return p
}
