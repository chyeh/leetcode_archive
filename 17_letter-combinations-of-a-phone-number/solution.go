package solution

func generate(sol string, digits string, alphabets map[byte]string, sols *[]string) {
	if len(digits) == 0 {
		*sols = append(*sols, sol)
		return
	}
	for _, c := range alphabets[digits[0]] {
		generate(sol+string(c), digits[1:], alphabets, sols)
	}
}

func letterCombinations(digits string) []string {
	sols := []string{}
	if len(digits) == 0 {
		return sols
	}
	alphabets := make(map[byte]string)
	alphabets['2'] = "abc"
	alphabets['3'] = "def"
	alphabets['4'] = "ghi"
	alphabets['5'] = "jkl"
	alphabets['6'] = "mno"
	alphabets['7'] = "pqrs"
	alphabets['8'] = "tuv"
	alphabets['9'] = "wxyz"
	generate("", digits, alphabets, &sols)
	return sols
}
