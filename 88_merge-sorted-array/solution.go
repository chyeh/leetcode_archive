package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	ans := make([]int, m+n, m+n)
	j := m - 1
	k := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if j == -1 {
			ans[i] = nums2[k]
			k--
			continue
		}
		if k == -1 {
			ans[i] = nums1[j]
			j--
			continue
		}
		if nums1[j] > nums2[k] {
			ans[i] = nums1[j]
			j--
		} else {
			ans[i] = nums2[k]
			k--
		}
	}
	copy(nums1, ans)
}
