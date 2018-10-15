package solution

// Solution I: Use map. Time Complexity O(n), Space Complexity O(n)
/*
func majorityElement(nums []int) int {
    cnt := make(map[int]int)
    for _, num := range nums {
        if v, ok := cnt[num]; !ok {
            cnt[num] = 1
        } else {
            cnt[num] = v + 1
        }
        if cnt[num] > len(nums) / 2 {
            return num
        }
    }
    return -1
}
*/

// Solution II:  Boyer-Moore Voting Algorithm. Time Complexity O(n), Space Complexity O(1)
func majorityElement(nums []int) int {
	var cnt, candidate int
	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}
		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}
	return candidate
}
