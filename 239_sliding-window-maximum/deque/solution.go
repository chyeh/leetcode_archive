package solution

type deque struct {
	data []int
}

func newDeque() *deque {
	return &deque{
		data: make([]int, 0),
	}
}

func (dq *deque) push(d int) {
	dq.data = append(dq.data, d)
}

func (dq *deque) popBack() int {
	tail := dq.data[len(dq.data)-1]
	dq.data = dq.data[:len(dq.data)-1]
	return tail
}

func (dq *deque) back() int {
	return dq.data[len(dq.data)-1]
}

func (dq *deque) popFront() int {
	head := dq.data[0]
	dq.data = dq.data[1:]
	return head
}

func (dq *deque) front() int {
	return dq.data[0]
}

func (dq *deque) isEmpty() bool {
	return len(dq.data) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1)
	dq := newDeque()
	for i := 0; i < len(nums); i++ {
		if !dq.isEmpty() && dq.front() <= i-k {
			dq.popFront()
		}
		for !dq.isEmpty() && nums[i] > nums[dq.back()] {
			dq.popBack()
		}
		dq.push(i)
		if i >= k-1 {
			ans = append(ans, nums[dq.front()])
		}
	}
	return ans
}
