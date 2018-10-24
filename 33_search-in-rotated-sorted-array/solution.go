package solution

// Solution I
/*
func binarySearch(nums []int, l, r, target int) int {
	for l <= r {
		m := l + ((r - l) / 2)
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func findPivot(nums []int) int {
	first := nums[0]
	l, r := 0, len(nums)-1
	for l < r {
		m := l + ((r - l) / 2)
		if nums[m] < first {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func search(nums []int, target int) int {
	if len(nums) < 1 || (len(nums) == 1 && nums[0] != target) {
		return -1
	}
	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	l, r := 0, len(nums)-1
	if nums[l] < nums[r] {
		return binarySearch(nums, l, r, target)
	}
	pivot := findPivot(nums)
	if nums[pivot] <= target && target <= nums[r] {
		return binarySearch(nums, pivot, r, target)
	}
	return binarySearch(nums, l, pivot-1, target)
}
*/

// Solution II: One Pass
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + ((r - l) / 2)
		if nums[m] == target {
			return m
		}
		if nums[m] < nums[r] {
			if nums[m] <= target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[l] <= target && target <= nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
