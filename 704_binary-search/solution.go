package solution

// Solution I: Iteration
func search(nums []int, target int) int {
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
	return -1
}

// Solution II: Recursion
/*
func find(nums []int, l, r, target int) int {
    if l > r {
        return -1
    }
    m := l + (r-l)/2
    if target == nums[m] {
        return m
    }
    if target < nums[m] {
        return find(nums, l, m-1, target)
    }
    return find(nums, m+1, r, target)
}

func search(nums []int, target int) int {
    return find(nums, 0, len(nums)-1, target)
}
*/
