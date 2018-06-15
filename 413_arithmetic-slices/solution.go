package solution

func isArithmatic(A []int, i int, j int) bool {
	d := (A[j] - A[i]) / (j - i)
	for k := i; k < j; k++ {
		if A[k+1]-A[k] != d {
			return false
		}
	}
	return true
}

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}

	var ans int
	for p := 0; p < len(A)-2; p++ {
		for q := p + 2; q < len(A); q++ {
			if isArithmatic(A, p, q) {
				ans++
			}
		}
	}

	return ans
}
