package solution

var res []int

func recur(n int) int {
	if n == 0 {
		return 0
	}
	if res[n] != 0 {
		return res[n]
	}
	return n%2 + recur(n/2)
}

func countBits(num int) []int {
	res = make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = recur(i)
	}
	return res
}
