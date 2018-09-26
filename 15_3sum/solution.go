package solution

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		valueToIndex := make(map[int]int)
		sol := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			if _, ok := sol[nums[j]]; ok {
				continue
			}
			if index, ok := valueToIndex[(-nums[i])-nums[j]]; ok {
				ans = append(ans, []int{nums[i], nums[index], nums[j]})
				sol[nums[j]] = true
			}
			if _, ok := valueToIndex[nums[j]]; !ok {
				valueToIndex[nums[j]] = j
			}
		}
	}
	return ans
}
