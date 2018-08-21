package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	for i, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = true
		if i >= k {
			delete(m, nums[i-k])
		}
	}
	return false
}
