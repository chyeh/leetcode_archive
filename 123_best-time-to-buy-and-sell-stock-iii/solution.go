package solution

// Solution I: Divide to Two Sections. Time Complexity: O(n^2)
/*
func maxProfitRange(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := -2147483648
	var maxDiff int
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		} else {
			if d := max - prices[i]; d > maxDiff {
				maxDiff = d
			}
		}
	}
	return maxDiff
}

func maxProfit(prices []int) int {
	var max int
	for i := range prices {
		p := maxProfitRange(prices[:i]) + maxProfitRange(prices[i:])
		if p > max {
			max = p
		}
	}
	return max
}
*/

// Solution II: Improved from Solution I. Time Complexity = O(n); Space Complexity = O(n)
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buyAfter := make([]int, len(prices))
	currMax := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > currMax {
			currMax = prices[i]
		}
		p := currMax - prices[i]
		if p > buyAfter[i+1] {
			buyAfter[i] = p
		} else {
			buyAfter[i] = buyAfter[i+1]
		}
	}

	sellBefore := make([]int, len(prices))
	currMin := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < currMin {
			currMin = prices[i]
		}
		p := prices[i] - currMin
		if p > sellBefore[i-1] {
			sellBefore[i] = p
		} else {
			sellBefore[i] = sellBefore[i-1]
		}
	}

	var max int
	for i := range prices {
		sum := sellBefore[i] + buyAfter[i]
		if sum > max {
			max = sum
		}
	}
	return max
}
