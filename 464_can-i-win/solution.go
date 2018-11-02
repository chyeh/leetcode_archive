package solution

// Time Limit Exceeded
/*
func canWin(choices []int, total int) bool {
	for i := range choices {
		if choices[i] >= total {
			return true
		}
		nextChoices := make([]int, len(choices))
		copy(nextChoices, choices)
		if !canWin(append(nextChoices[:i], nextChoices[i+1:]...), total-choices[i]) {
			return true
		}
	}
	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if (maxChoosableInteger+1)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	choices := make([]int, maxChoosableInteger)
	for i := range choices {
		choices[i] = i + 1
	}
	return canWin(choices, desiredTotal)
}
*/

// Solution I: Top-down Dynamic Programming
func canWin(choices [21]bool, results map[[21]bool]bool, sum, total int) bool {
	if v, ok := results[choices]; ok {
		return v
	}
	for i := range choices {
		if !choices[i] {
			continue
		}
		choices[i] = false
		if i+sum >= total {
			results[choices] = true
			return true
		}
		if !canWin(choices, results, i+sum, total) {
			choices[i] = true
			results[choices] = true
			return true
		}
		choices[i] = true
	}
	results[choices] = false
	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if (maxChoosableInteger+1)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	var choices [21]bool
	for i := 1; i <= maxChoosableInteger; i++ {
		choices[i] = true
	}
	results := make(map[[21]bool]bool)
	return canWin(choices, results, 0, desiredTotal)
}
