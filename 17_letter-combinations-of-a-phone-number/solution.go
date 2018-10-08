package solution

func generate(curr string, sets []string, ans *[]string) {
	if len(curr) == len(sets) {
		*ans = append(*ans, curr)
		return
	}
	for _, c := range sets[len(curr)] {
		generate(curr+string(c), sets, ans)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	alphabets := make(map[rune]string)
	alphabets['2'] = "abc"
	alphabets['3'] = "def"
	alphabets['4'] = "ghi"
	alphabets['5'] = "jkl"
	alphabets['6'] = "mno"
	alphabets['7'] = "pqrs"
	alphabets['8'] = "tuv"
	alphabets['9'] = "wxyz"
	sets := []string{}
	for _, d := range digits {
		sets = append(sets, alphabets[d])
	}
	ans := []string{}
	generate("", sets, &ans)
	return ans
}
