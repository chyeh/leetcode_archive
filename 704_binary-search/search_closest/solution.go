package solution

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	if target > nums[len(nums)-1] {
		return nums[len(nums)-1]
	}
	if target < nums[0] {
		return nums[0]
	}

	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		}
		if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if nums[l]-target < target-nums[r] {
		return nums[l]
	}
	return nums[r]
}
