package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := (len(nums1) + len(nums2))
	merged := merge(nums1, nums2)
	if l%2 == 0 {
		return (float64(merged[l/2]) + float64(merged[l/2-1])) / 2
	}
	return float64(merged[l/2])
}

func merge(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil {
		return nil
	}
	if len(nums1) == 0 && len(nums2) == 0 {
		return []int{}
	}
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	res := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}

		if i == len(nums1) {
			res = append(res, nums2[j:]...)
			break
		}
		if j == len(nums2) {
			res = append(res, nums1[i:]...)
			break
		}
	}
	return res
}
