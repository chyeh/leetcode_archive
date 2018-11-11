package solution

// Time Limit Exceeded
/*
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	var rectangle int
	for currHeight := 1; ; currHeight++ {
		var maxWidth int
		l, r := 0, 0
		for ; r < len(heights); r++ {
			if heights[r] < currHeight {
				l = r + 1
				continue
			}
			maxWidth = max(maxWidth, r-l+1)
		}
		if maxWidth == 0 {
			break
		}
		rectangle = max(rectangle, currHeight*maxWidth)
	}
	return rectangle
}
*/

// Solution I: Time complexity O(N^2); Space complexity O(N).
/*
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	var rectangle int
	for _, currHeight := range heights {
		var maxWidth int
		l, r := 0, 0
		for ; r < len(heights); r++ {
			if heights[r] < currHeight {
				l = r + 1
				continue
			}
			maxWidth = max(maxWidth, r-l+1)
		}
		rectangle = max(rectangle, currHeight*maxWidth)
	}
	return rectangle
}
*/

// Solution II: Time complexity O(N^2); Space complexity O(1)
/*
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	var rectangle int
	for i, currHeight := range heights {
		l, r := i, i
		for ; r < len(heights) && heights[r] >= currHeight; r++ {
		}
		for ; l >= 0 && heights[l] >= currHeight; l-- {
		}
		width := (r - l) - 1
		rectangle = max(rectangle, currHeight*width)
	}
	return rectangle
}
*/

// Solution III: Stack. Time complexity O(N); Space complexity O(N).
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

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	s := newStack()
	var rectangle int
	for i, currHeight := range heights {
		for !s.isEmpty() && currHeight <= heights[s.top()] {
			curr := s.pop()
			h := heights[curr]
			var w int
			if s.isEmpty() {
				w = i
			} else {
				w = i - s.top() - 1
			}
			rectangle = max(rectangle, h*w)
		}
		s.push(i)
	}
	return rectangle
}
