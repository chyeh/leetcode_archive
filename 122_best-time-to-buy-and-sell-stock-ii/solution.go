package solution

func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		d := prices[i] - prices[i-1]
		if d > 0 {
			profit += d
		}
	}
	return profit
}
