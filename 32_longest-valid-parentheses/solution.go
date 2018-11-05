package solution

// Solution I: Dynamic Programming
/*
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	max := 0
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
			continue
		}

		startIndex := i - dp[i-1] - 1
		if startIndex >= 0 && s[startIndex] == '(' {
			dp[i] = dp[i-1] + 2
		} else {
			dp[i] = 0
			continue
		}

		prevIndex := startIndex - 1
		if prevIndex >= 0 {
			dp[i] += dp[prevIndex]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
*/

// Solution II: Stack
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

func longestValidParentheses(s string) int {
	var max int
	st := newStack()
	st.push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if st.top() == -1 || s[st.top()] == ')' {
				st.push(i)
			} else {
				st.pop()
				if length := i - st.top(); length > max {
					max = length
				}
			}
		} else {
			st.push(i)
		}
	}
	return max
}
