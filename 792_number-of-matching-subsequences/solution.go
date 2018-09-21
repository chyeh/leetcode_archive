package solution

func isSubseq(m map[rune][]int, word string) bool {
	currIndex := -1
	for _, c := range word {
		indexArr, ok := m[c]
		if !ok {
			return false
		}
		i := 0
		for ; i < len(indexArr); i++ {
			if indexArr[i] > currIndex {
				currIndex = indexArr[i]
				break
			}
		}
		if i == len(indexArr) {
			return false
		}
	}
	return true
}

func numMatchingSubseq(S string, words []string) int {
	indexMap := make(map[rune][]int, 0)
	for i, c := range S {
		if _, ok := indexMap[c]; !ok {
			indexMap[c] = []int{i}
		} else {
			indexMap[c] = append(indexMap[c], i)
		}
	}

	var cnt int
	for _, word := range words {
		if isSubseq(indexMap, word) {
			cnt++
		}
	}
	return cnt
}
