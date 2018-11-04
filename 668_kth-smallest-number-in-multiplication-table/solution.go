package solution

// Time Limit Exceeded
/*
func findKthNumber(m int, n int, k int) int {
	ans := make([]int, m*n+1)
	for i := 0; i <= m*n+1; i++ {
		ans = append(ans, 0)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j*i > k {
				break
			}
			ans[j*i]++
		}
	}

	for i := 1; i < len(ans); i++ {
		k = k - ans[i]
		if k <= 0 {
			return i
		}
	}
	return -1
}
*/

// Solution
func numOfElementsSmallerOrEqualTo(m, n, target int) int {
	var cnt int
	i := 1
	j := n
	for i <= m && j >= 1 {
		if target >= i*j {
			cnt += j
			i++
		} else {
			j--
		}
	}
	return cnt
}

func findKthNumber(m int, n int, k int) int {
	low, high := 1, m*n
	for low < high {
		target := low + (high-low)/2
		cnt := numOfElementsSmallerOrEqualTo(m, n, target)
		if cnt >= k {
			high = target
		} else {
			low = target + 1
		}
	}
	return low
}
