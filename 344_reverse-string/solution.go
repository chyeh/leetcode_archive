package solution

func reverseString(s string) string {
	slc := []rune(s)
	for l, r := 0, len(slc)-1; l < r; l, r = l+1, r-1 {
		slc[l], slc[r] = slc[r], slc[l]
	}
	return string(slc)
}
