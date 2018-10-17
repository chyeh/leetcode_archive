package solution

// Solution I: Brute Force
/*
func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	chars := make(map[rune]int)
	for _, c := range a {
		if v, ok := chars[c]; ok {
			chars[c] = v + 1
		} else {
			chars[c] = 1
		}
	}
	for _, c := range b {
		if v, ok := chars[c]; !ok {
			return false
		} else if v == 0 {
			return false
		} else {
			chars[c] = v - 1
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	var sols [][]string
	for len(strs) != 0 {
		sol := []string{strs[0]}
		strs = strs[1:]
		for j := 0; j < len(strs); {
			if isAnagram(sol[0], strs[j]) {
				sol = append(sol, strs[j])
				strs = append(strs[:j], strs[j+1:]...)
			} else {
				j++
			}
		}
		sols = append(sols, sol)
	}
	return sols
}
*/

// Solution II: Categorization by Sorting
// Skip...

// Solution III: Categorization by Counting
func groupAnagrams(strs []string) [][]string {
	cntToStr := make(map[[26]int][]string)
	for _, str := range strs {
		var cnt [26]int
		for _, c := range str {
			cnt[c-'a']++
		}
		if _, ok := cntToStr[cnt]; !ok {
			cntToStr[cnt] = []string{str}
		} else {
			cntToStr[cnt] = append(cntToStr[cnt], str)
		}
	}
	var sols [][]string
	for _, v := range cntToStr {
		sols = append(sols, v)
	}
	return sols
}
