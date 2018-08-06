package solution

import (
	"math/rand"
)

type RandomizedSet struct {
	valToInd map[int]int
	valSlice []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		make(map[int]int),
		make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToInd[val]; ok {
		return false
	}
	this.valToInd[val] = len(this.valSlice)
	this.valSlice = append(this.valSlice, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.valToInd[val]; !ok {
		return false
	}
	i := this.valToInd[val]

	// swap
	tmp := this.valSlice[i]
	this.valSlice[i] = this.valSlice[len(this.valSlice)-1]
	this.valSlice[len(this.valSlice)-1] = tmp
	// update map
	this.valToInd[val] = len(this.valSlice) - 1
	this.valToInd[this.valSlice[i]] = i

	delete(this.valToInd, val)
	this.valSlice = this.valSlice[:len(this.valSlice)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.valSlice))
	return this.valSlice[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
