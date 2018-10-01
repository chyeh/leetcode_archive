package solution

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

func (s *stack) push(d int) {
	s.data = append(s.data, d)
}

func (s *stack) pop() int {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *stack) top() int {
	return s.data[len(s.data)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func sumSubarrayMins(A []int) int {
	length := len(A)
	var sum int
	A = append(A, -2147483648)
	s := newStack()
	for r := 0; r <= length; r++ {
		for !s.isEmpty() && s.top() > A[r] {
			m := s.pop()
			l := -1
			if !s.isEmpty() {
				l = s.top()
				sum += A[m] * (m - l) * (r - m)
			}
		}
		if r != length {
			s.push(r)
		}
	}
	return sum % (1000000007)
}
