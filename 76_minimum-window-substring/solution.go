package solution

func minWindow(s string, t string) string {
	targetNum := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		if _, ok := targetNum[t[i]]; !ok {
			targetNum[t[i]] = 1
		} else {
			targetNum[t[i]]++
		}
	}
	min := 2147483647
	minh, mint := 0, 0
	head, tail, cnt := 0, 0, 0
	for ; tail < len(s); tail++ {
		if _, ok := targetNum[s[tail]]; !ok {
			continue
		}
		if targetNum[s[tail]] > 0 {
			cnt++
		}
		targetNum[s[tail]]--

		for cnt == len(t) {
			for ; ; head++ {
				if _, ok := targetNum[s[head]]; ok {
					break
				}
			}
			if targetNum[s[head]] >= 0 {
				cnt--
			}
			targetNum[s[head]]++

			length := tail - head + 1
			if length < min {
				min = length
				minh = head
				mint = tail
			}
			head++
		}
	}

	if min == 2147483647 {
		return ""
	}
	return s[minh : mint+1]
}
