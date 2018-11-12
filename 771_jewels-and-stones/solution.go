package solution

func numJewelsInStones(J string, S string) int {
	jewelry := make(map[byte]bool)
	for i := range J {
		jewelry[J[i]] = true
	}
	var cnt int
	for i := range S {
		if _, ok := jewelry[S[i]]; ok {
			cnt++
		}
	}
	return cnt
}
