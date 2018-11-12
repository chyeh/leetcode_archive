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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxRectangle(height []int) int {
	var rectangle int
	height = append(height, 0)
	s := newStack()
	for i, h := range height {
		for !s.isEmpty() && height[s.top()] >= h {
			curr := s.pop()
			var w int
			if s.isEmpty() {
				w = i
			} else {
				w = i - s.top() - 1
			}
			rectangle = max(rectangle, height[curr]*w)
		}
		s.push(i)
	}
	return rectangle
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) < 1 {
		return 0
	}
	height := make([]int, len(matrix[0]))
	var rectangle int
	for i := 0; i < len(matrix); i++ {
		if i == 0 {
			for i := range matrix[0] {
				if matrix[0][i] == '0' {
					height[i] = 0
				} else {
					height[i] = 1
				}
			}
		} else {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == '0' {
					height[j] = 0
				} else {
					height[j] = height[j] + 1
				}
			}
		}
		rectangle = max(rectangle, findMaxRectangle(height))
	}
	return rectangle
}
