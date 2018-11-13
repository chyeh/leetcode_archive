package solution

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func topKFrequent(nums []int, k int) []int {
	countArr := make([][]int, len(nums)+1)
	for i := range countArr {
		countArr[i] = make([]int, 0)
	}
	countOfNum := make(map[int]int)
	for _, num := range nums {
		if v, ok := countOfNum[num]; ok {
			countOfNum[num] = v + 1
		} else {
			countOfNum[num] = 1
		}
	}

	for k, v := range countOfNum {
		countArr[v] = append(countArr[v], k)
	}
	ans := make([]int, 0, k)
	for i, cnt := len(countArr)-1, 0; i >= 0 && cnt < k; i-- {
		if len(countArr[i]) != 0 {
			length := min(len(countArr[i]), k-cnt)
			ans = append(ans, countArr[i][:length]...)
			cnt = cnt + length
		}
	}
	return ans
}
