package solution

func isHappy(n int) bool {
	set := make(map[int]struct{})
	for i := n; i != 1; i = nextNum(i) {
		if _, ok := set[i]; ok {
			return false
		}
		set[i] = struct{}{}
	}
	return true
}

func nextNum(n int) int {
	ans := 0
	for ; n != 0; n /= 10 {
		d := n % 10
		ans += d * d
	}
	return ans
}
