package solution

func moveZeroes(nums []int) {
	for i, zeroInd := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// if nums[zeroInd] is already non-zero, order remains the same
			tmp := nums[i]
			nums[i] = nums[zeroInd]
			nums[zeroInd] = tmp
			zeroInd++
		}
	}
}
