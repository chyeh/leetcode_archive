package solution

// Solution I: Time complexity O(lgN+K)
func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	} else if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		if arr[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	var index int
	if arr[l]-x > arr[l-1]-x {
		index = l - 1
	} else {
		index = l
	}
	head, tail := index, index+1
	for tail-head < k {
		if head == 0 {
			tail++
			continue
		}
		if tail == len(arr) {
			head--
			continue
		}
		if x-arr[head-1] <= arr[tail]-x {
			head--
		} else {
			tail++
		}
	}
	return arr[head:tail]
}

// Solution II: O(lgN)
// Pending...
