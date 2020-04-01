package solution

func singleNumber(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := set[num]; ok {
			delete(set, num)
		} else {
			set[num] = struct{}{}
		}
	}
	var ans int
	for ans = range set {
	}
	return ans
}
