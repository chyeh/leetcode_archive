package solution

import (
	"math/rand"
)

type arraylist struct {
	data       []int
	dataTondex map[int][]int
}

func NewArrayList() *arraylist {
	return &arraylist{
		data:       make([]int, 0),
		dataTondex: make(map[int][]int),
	}
}

func (l *arraylist) pushback(d int) {
	l.data = append(l.data, d)
	if _, ok := l.dataTondex[d]; ok {
		l.dataTondex[d] = append(l.dataTondex[d], len(l.data)-1)
	} else {
		l.dataTondex[d] = []int{len(l.data) - 1}
	}
}

func (l *arraylist) popback() int {
	tail := l.data[len(l.data)-1]
	l.data = l.data[:len(l.data)-1]
	l.dataTondex[tail] = l.dataTondex[tail][:len(l.dataTondex[tail])-1]
	return tail
}

func (l *arraylist) remove(d int) {
	// must have d in the array list
	index := l.dataTondex[d][len(l.dataTondex[d])-1]
	l.data = append(l.data[:index], l.data[index+1:]...)
	l.dataTondex[d] = l.dataTondex[d][:len(l.dataTondex[d])-1]
}

func (l *arraylist) isEmpty() bool {
	return len(l.data) == 0
}

type RandomizedCollection struct {
	valToInd map[int]*arraylist
	valSlice []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		make(map[int]*arraylist),
		make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	if _, ok := this.valToInd[val]; ok {
		this.valToInd[val].pushback(len(this.valSlice))
		this.valSlice = append(this.valSlice, val)
		return false
	}
	this.valToInd[val] = NewArrayList()
	this.valToInd[val].pushback(len(this.valSlice))
	this.valSlice = append(this.valSlice, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if _, ok := this.valToInd[val]; !ok {
		return false
	}

	top := this.valToInd[val].popback()

	if top != len(this.valSlice)-1 {
		// overwrite
		tail := this.valSlice[len(this.valSlice)-1]
		this.valSlice[top] = tail
		// update map
		this.valToInd[tail].remove(len(this.valSlice) - 1)
		this.valToInd[tail].pushback(top)
	}

	if this.valToInd[val].isEmpty() {
		delete(this.valToInd, val)
	}
	this.valSlice = this.valSlice[:len(this.valSlice)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedCollection) GetRandom() int {
	i := rand.Intn(len(this.valSlice))
	return this.valSlice[i]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
