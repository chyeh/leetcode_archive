package solution

func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 1
	var cnt, sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := sumMap[sum-k]; ok {
			cnt = cnt + v
		}
		if v, ok := sumMap[sum]; ok {
			sumMap[sum] = v + 1
		} else {
			sumMap[sum] = 1
		}
	}
	return cnt
}

/** Cummulative Sum:
 *  Time: O(N^)
 *  Space: O(1)
func subarraySum(nums []int, k int) int {
	var cnt int
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}
*/
