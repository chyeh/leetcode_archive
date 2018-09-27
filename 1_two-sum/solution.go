package solution

func twoSum(nums []int, target int) []int {
	valueToIndex := make(map[int]int)
	for i, num := range nums {
		if index, ok := valueToIndex[target-num]; ok {
			return []int{index, i}
		}
		valueToIndex[num] = i
	}
	return nil
}
