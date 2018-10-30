package solution

// Solution I: Brute Force
/*
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxArea(height []int) int {
    var ans int
    for i := 0; i < len(height) - 1; i++ {
        for j := i + 1; j < len(height); j++ {
            h := min(height[i], height[j])
            w := j-i
            ans = max(ans, h*w)
        }
    }
    return ans
}
*/

// Solution II: Two Pointer
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	var ans int
	for l, r := 0, len(height)-1; l < r; {
		water := (r - l) * min(height[r], height[l])
		ans = max(ans, water)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
