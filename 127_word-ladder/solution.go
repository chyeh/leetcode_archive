package solution

import (
	"container/list"
)

type queue struct {
	*list.List
}

func NewQueue() *queue {
	return &queue{
		list.New(),
	}
}

func (q *queue) push(d interface{}) {
	q.PushBack(d)
}

func (q *queue) pop() interface{} {
	return q.Remove(q.Front())
}

func (q *queue) isEmpty() bool {
	return q.Len() == 0
}

func isValid(next, word string) bool {
	cnt := 0
	for i, _ := range word {
		if next[i] != word[i] {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append([]string{beginWord}, wordList...)
	distance := make([]int, len(wordList))
	visited := make([]bool, len(wordList))
	q := NewQueue()
	for i, w := range wordList {
		if w == endWord {
			distance[i] = 1
			visited[i] = true
			q.push(i)
		}
	}
	if q.isEmpty() {
		return 0
	}

	for !q.isEmpty() {
		popped := q.pop().(int)
		for i, next := range wordList {
			if !visited[i] && isValid(next, wordList[popped]) {
				visited[i] = true
				distance[i] = distance[popped] + 1
				q.push(i)
			}
		}
	}
	return distance[0]
}
