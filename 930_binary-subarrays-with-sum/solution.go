package solution

// Solution I: Prefix Sum. Time Complexity O(N^2); Space Complexity O(N).
/*
func numSubarraysWithSum(A []int, S int) int {
	p := make([]int, len(A)+1)
	for i := 1; i < len(p); i++ {
		p[i] = p[i-1] + A[i-1]
	}
	var cnt int
	for i := 0; i < len(A); i++ {
		for j := i + 1; j <= len(A); j++ {
			if p[j]-p[i] == S {
				cnt++
			}
		}
	}
	return cnt
}
*/

// Solution II: Improved Prefix Sum. Time Complexity O(N); Space Complexity O(N).
/*
func numSubarraysWithSum(A []int, S int) int {
	p := make([]int, len(A)+1)
	for i := 1; i < len(p); i++ {
		p[i] = p[i-1] + A[i-1]
	}
	var cnt int
	m := make(map[int]int)
	for _, psum := range p {
		if v, ok := m[psum]; ok {
			cnt += v
		}
		if v, ok := m[psum+S]; ok {
			m[psum+S] = v + 1
		} else {
			m[psum+S] = 1
		}
	}
	return cnt
}
*/

// Solution III: Three Pointers. Time Complexity O(N); Space Complexity O(1).
// Pending...
