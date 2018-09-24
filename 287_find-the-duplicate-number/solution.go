package solution

func findDuplicate(nums []int) int {
	tortoise, hare := nums[0], nums[0]
	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}
	i, j := nums[0], hare
	for i != j {
		i = nums[i]
		j = nums[j]
	}
	return i
}
