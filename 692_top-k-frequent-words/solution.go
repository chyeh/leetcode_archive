package solution

func heapify(freq map[string]int, heap []string, heapIndex map[string]int, index int) {
	largestIndex := index
	l := (2 * index) + 1
	r := (2 * index) + 2
	if l < len(heap) &&
		((freq[heap[l]] > freq[heap[largestIndex]]) || (freq[heap[l]] == freq[heap[largestIndex]] && heap[l] < heap[largestIndex])) {
		largestIndex = l
	}
	if r < len(heap) &&
		((freq[heap[r]] > freq[heap[largestIndex]]) || (freq[heap[r]] == freq[heap[largestIndex]] && heap[r] < heap[largestIndex])) {
		largestIndex = r
	}
	if largestIndex != index {
		heap[largestIndex], heap[index] = heap[index], heap[largestIndex]
		heapIndex[heap[largestIndex]], heapIndex[heap[index]] = largestIndex, index
		heapify(freq, heap, heapIndex, largestIndex)
	}
}

func heapIncreasKey(freq map[string]int, heap []string, heapIndex map[string]int, word string) {
	freq[word]++
	index := heapIndex[word]
	parent := (index - 1) / 2
	for freq[heap[parent]] < freq[heap[index]] || (freq[heap[parent]] == freq[heap[index]] && heap[parent] > heap[index]) {
		heap[parent], heap[index] = heap[index], heap[parent]
		heapIndex[heap[parent]], heapIndex[heap[index]] = parent, index
		index = (index - 1) / 2
		parent = (parent - 1) / 2
	}
}

func extractMax(freq map[string]int, heap *[]string, heapIndex map[string]int) string {
	max := (*heap)[0]
	(*heap)[0], (*heap)[len((*heap))-1] = (*heap)[len((*heap))-1], (*heap)[0]
	heapIndex[(*heap)[0]], heapIndex[(*heap)[len((*heap))-1]] = 0, len((*heap))-1
	(*heap) = (*heap)[:len((*heap))-1]
	heapify(freq, (*heap), heapIndex, 0)
	return max
}

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int, len(words))
	var heap []string
	heapIndex := make(map[string]int, len(words))
	for _, w := range words {
		if _, ok := freq[w]; !ok {
			heap = append(heap, w)
			heapIndex[w] = len(heap) - 1
		}
		heapIncreasKey(freq, heap, heapIndex, w)
	}
	var ans []string
	for i := 0; i < k; i++ {
		ans = append(ans, extractMax(freq, &heap, heapIndex))
	}
	return ans
}
