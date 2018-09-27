package solution

func reverse(s []rune) []rune {
	slc := []rune(s)
	for l, r := 0, len(slc)-1; l < r; l, r = l+1, r-1 {
		slc[l], slc[r] = slc[r], slc[l]
	}
	return slc
}

func addBinary(a string, b string) string {
	aSlc := reverse([]rune(a))
	bSlc := reverse([]rune(b))
	if len(a) > len(b) {
		for i := 0; i < len(a)-len(b); i++ {
			bSlc = append(bSlc, '0')
		}
	} else if len(a) < len(b) {
		for i := 0; i < len(b)-len(a); i++ {
			aSlc = append(aSlc, '0')
		}
	}
	carry := '0'
	var ans []rune
	for i, j := 0, 0; i < len(aSlc) && j < len(bSlc); i, j = i+1, j+1 {
		if aSlc[i] == '1' && bSlc[j] == '1' {
			ans = append(ans, carry)
			carry = '1'
		} else if aSlc[i] == '0' && bSlc[j] == '0' {
			ans = append(ans, carry)
			carry = '0'
		} else {
			if carry == '1' {
				ans = append(ans, '0')
				carry = '1'
			} else {
				ans = append(ans, '1')
				carry = '0'
			}
		}
	}
	if carry == '1' {
		ans = append(ans, '1')
	}
	return string(reverse(ans))
}
