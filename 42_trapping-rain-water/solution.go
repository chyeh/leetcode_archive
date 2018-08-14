package solution

func trap(height []int) int {
	level := make([]int, len(height))
	for i := 0; i < len(level); i++ {
		level[i] = -1

		rhighest := height[i]
		ri := -1
		for j := i + 1; j < len(height); j++ {
			if height[j] > rhighest {
				rhighest = height[j]
				ri = j
			}
		}
		lhighest := height[i]
		li := -1
		for j := i - 1; j >= 0; j-- {
			if height[j] > lhighest {
				lhighest = height[j]
				li = j
			}
		}
		if ri != -1 && li != -1 {
			if height[ri] < height[li] {
				level[i] = height[ri]
			} else {
				level[i] = height[li]
			}
		}
	}

	ans := 0
	for i, h := range height {
		if level[i] != -1 {
			ans = ans + (level[i] - h)
		}
	}
	return ans
}
