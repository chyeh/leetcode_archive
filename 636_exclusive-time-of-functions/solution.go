package solution

import (
	"strconv"
	"strings"
)

type stack struct {
	data []int
}

func NewStack() *stack {
	return &stack{
		data: []int{},
	}
}

func (s *stack) push(id int) {
	s.data = append(s.data, id)
}

func (s *stack) pop() int {
	// panic if stack is empty
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) top() int {
	if s.isEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}

func parseLog(log string) (int, string, int) {
	parsed := strings.Split(log, ":")
	var funcId, timestamp int64
	var op string
	var err error
	if funcId, err = strconv.ParseInt(parsed[0], 10, 64); err != nil {
		panic(err)
	}
	op = parsed[1]
	if timestamp, err = strconv.ParseInt(parsed[2], 10, 64); err != nil {
		panic(err)
	}
	return int(funcId), op, int(timestamp)
}

func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	cnt := 0
	s := NewStack()
	for _, log := range logs {
		id, op, time := parseLog(log)
		timeDiff := time - cnt
		cnt = time
		if op == "start" {
			if !s.isEmpty() {
				top := s.top()
				ans[top] += timeDiff
			}
			s.push(id)
		} else { // stop
			top := s.top()
			ans[top] += (timeDiff + 1)
			cnt++
			s.pop()
		}
	}
	return ans
}
