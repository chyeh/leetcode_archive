package solution

type PriorityQueue struct {
	heap []string
	freq map[string]int
}

func NewPriorityQueue(freq map[string]int) *PriorityQueue {
	pq := &PriorityQueue{
		freq: freq,
	}
	for word := range freq {
		pq.Insert(word)
	}
	return pq
}

func (q *PriorityQueue) Insert(word string) {
	q.heap = append(q.heap, word)
	index := len(q.heap) - 1
	parent := (index - 1) / 2
	for index > 0 && q.greaterThan(q.heap[index], q.heap[parent]) {
		q.heap[index], q.heap[parent] = q.heap[parent], q.heap[index]
		index = (index - 1) / 2
		parent = (parent - 1) / 2
	}
}

func (q PriorityQueue) greaterThan(a, b string) bool {
	if q.freq[a] > q.freq[b] {
		return true
	} else if q.freq[a] == q.freq[b] {
		if a < b {
			return true
		}
	}
	return false
}

func (q *PriorityQueue) ExtractMax() string {
	max := q.heap[0]
	q.heap[0], q.heap[len(q.heap)-1] = q.heap[len(q.heap)-1], q.heap[0]
	q.heap = q.heap[:len(q.heap)-1]
	q.heapify(0)
	return max
}

func (q *PriorityQueue) heapify(index int) {
	largest := index
	l := (2 * index) + 1
	r := (2 * index) + 2
	if l < len(q.heap) && q.greaterThan(q.heap[l], q.heap[largest]) {
		largest = l
	}
	if r < len(q.heap) && q.greaterThan(q.heap[r], q.heap[largest]) {
		largest = r
	}
	if largest != index {
		q.heap[index], q.heap[largest] = q.heap[largest], q.heap[index]
		q.heapify(largest)
	}
}

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int, len(words))
	for _, w := range words {
		if _, ok := freq[w]; !ok {
			freq[w] = 1
		} else {
			freq[w]++
		}
	}
	pq := NewPriorityQueue(freq)

	var ans []string
	for i := 0; i < k; i++ {
		ans = append(ans, pq.ExtractMax())
	}
	return ans
}
