package solution

// Solution I: Descending Deque
/*
type deque struct {
	data []int
}

func newDeque() *deque {
	return &deque{
		data: make([]int, 0),
	}
}

func (s *deque) pushBack(d int) {
	s.data = append(s.data, d)
}

func (s *deque) pop() int {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *deque) back() int {
	return s.data[len(s.data)-1]
}

func (s *deque) front() int {
	return s.data[0]
}

func (s *deque) isEmpty() bool {
	return len(s.data) == 0
}

func maxProfit(prices []int) int {
	var max int
	s := newDeque()
	for i := len(prices) - 1; i >= 0; i-- {
		for !s.isEmpty() && s.back() <= prices[i] {
			s.pop()
		}
		if !s.isEmpty() {
			d := s.front() - prices[i]
			if d > max {
				max = d
			}
		}
		s.pushBack(prices[i])
	}
	return max
}
*/

// Solution II: Two Pointers
func maxProfit(prices []int) int {
	max := -2147483648
	var maxDiff int
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		} else {
			if d := max - prices[i]; d > maxDiff {
				maxDiff = d
			}
		}
	}
	return maxDiff
}
