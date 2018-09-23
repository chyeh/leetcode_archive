package solution

func maxChunksToSorted(arr []int) int {
	var max, cnt int
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if i == max {
			cnt++
		}
	}
	return cnt
}
