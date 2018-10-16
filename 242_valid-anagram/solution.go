package solution

// Solution I: Sort
/*
type RuneSlice []rune

func (rs RuneSlice) Less(i int, j int) bool { return rs[i] < rs[j] }
func (rs RuneSlice) Swap(i int, j int)      { rs[i], rs[j] = rs[j], rs[i] }
func (rs RuneSlice) Len() int               { return len(rs) }

func isAnagram(s string, t string) bool {
	sSlc := RuneSlice(s)
	tSlc := RuneSlice(t)
	sort.Sort(sSlc)
	sort.Sort(tSlc)
	return reflect.DeepEqual(sSlc, tSlc)
}
*/

// Solution II: Map
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make(map[rune]int)
	for _, c := range s {
		if v, ok := chars[c]; ok {
			chars[c] = v + 1
		} else {
			chars[c] = 1
		}
	}
	for _, c := range t {
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
