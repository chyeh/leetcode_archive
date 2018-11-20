package solution

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i := 0
	for i < len(A)-1 && A[i] < A[i+1] {
		i++
	}
	if i == len(A)-1 || i == 0 || A[i] == A[i+1] {
		return false
	}
	for i < len(A)-1 && A[i] > A[i+1] {
		i++
	}
	if i == len(A)-1 {
		return true
	}
	return false
}
