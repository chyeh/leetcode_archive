package solution

type queue struct {
	data []string
}

func newQueue() *queue {
	return &queue{
		data: make([]string, 0),
	}
}

func (q *queue) push(s string) {
	q.data = append(q.data, s)
}

func (q *queue) pop() string {
	head := q.data[0]
	q.data = q.data[1:]
	return head
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

func isValid(word, next string) bool {
	cnt := 0
	for i := range next {
		if word[i] != next[i] {
			cnt++
		}
	}
	if cnt > 1 {
		return false
	}
	return true
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	visited := make(map[string]bool)
	nextWords := make(map[string][]string)
	distance := make(map[string]int)
	for _, w := range wordList {
		visited[w] = false
		distance[w] = 0
	}
	if _, ok := visited[beginWord]; !ok {
		wordList = append([]string{beginWord}, wordList...)
		visited[beginWord] = false
	}
	if _, ok := visited[endWord]; !ok {
		return nil
	}
	q := newQueue()
	visited[endWord] = true
	q.push(endWord)

	for !q.isEmpty() {
		popped := q.pop()
		for _, next := range wordList {
			if !visited[next] && isValid(popped, next) {
				visited[next] = true
				q.push(next)
				distance[next] = distance[popped] + 1
				nextWords[next] = []string{popped}
			} else if isValid(popped, next) {
				if distance[next] > distance[popped]+1 {
					nextWords[next] = []string{popped}
				} else if distance[next] == distance[popped]+1 {
					nextWords[next] = append(nextWords[next], popped)
				}
			}
		}
	}
	if distance[beginWord] == 0 {
		return nil
	}
	sols := [][]string{{beginWord}}
	for step := 0; sols[0][step] != endWord; step++ {
		numSols := len(sols)
		for i := 0; i < numSols; i++ {
			if len(nextWords[sols[i][step]]) > 1 {
				words := nextWords[sols[i][step]]
				sols[i] = append(sols[i], words[0])
				for j := 1; j < len(words); j++ {
					copied := make([]string, len(sols[i]))
					copy(copied, sols[i])
					copied[step+1] = words[j]
					sols = append(sols, copied)
				}
			} else {
				word := nextWords[sols[i][step]][0]
				sols[i] = append(sols[i], word)
			}
		}
	}
	return sols
}

func dfs(begin, end string, wordList []string, sol []string, sols *[][]string, visited map[string]bool, length *int) {
	if len(sol) > *length {
		return
	}
	if begin == end {
		if len(sol) <= *length {
			*length = len(sol)
			d := make([]string, len(sol))
			copy(d, sol)
			*sols = append(*sols, d)
		}
		return
	}

	visited[begin] = true
	for _, word := range wordList {
		if !visited[word] && isValid(begin, word) {
			dfs(word, end, wordList, append(sol, word), sols, visited, length)
		}
	}
	visited[begin] = false
}
