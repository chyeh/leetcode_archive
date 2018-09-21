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

func (s stack) top() int {
	return s.data[len(s.data)-1]
}

func (s stack) isEmpty() bool {
	return len(s.data) == 0
}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	s := newStack()
	for i := len(temperatures) - 1; i >= 0; i-- {
		for !s.isEmpty() && temperatures[i] >= temperatures[s.top()] {
			s.pop()
		}
		if !s.isEmpty() {
			ans[i] = s.top() - i
		}
		s.push(i)
	}
	return ans
}
