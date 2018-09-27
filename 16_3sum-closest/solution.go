package solution

import (
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	minDiff := 2147483647
	var ans int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			d := abs(target - sum)
			if d == 0 {
				return sum
			}
			if d < minDiff {
				minDiff = d
				ans = sum
			}
			if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}
