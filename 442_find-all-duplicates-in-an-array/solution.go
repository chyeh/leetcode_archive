package solution

func findDuplicates(nums []int) []int {
	records := make(map[int]int)
	for _, n := range nums {
		if _, ok := records[n]; ok {
			records[n] = records[n] + 1
		} else {
			records[n] = 1
		}
	}

	var ans []int
	for key, r := range records {
		if r == 2 {
			ans = append(ans, key)
		}
	}
	return ans
}
