package solution

// Solution I: Use set.
/*
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
*/

// Solution II: No extra memory.
func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
