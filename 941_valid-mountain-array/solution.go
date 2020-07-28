package solution

func validMountainArray(A []int) bool {
	i := 0
	for ; i < len(A)-1 && A[i] < A[i+1]; i++ {
	}
	if i == 0 || i == len(A)-1 {
		return false
	}
	for ; i < len(A)-1 && A[i] > A[i+1]; i++ {
	}
	return i == len(A)-1
}
