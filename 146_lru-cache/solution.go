package solution

import (
	"container/list"
)

type LRUCache struct {
	kv       map[int]int
	kn       map[int]*list.Element
	kl       *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		kv:       make(map[int]int),
		kn:       make(map[int]*list.Element),
		kl:       list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.kv[key]; !ok {
		return -1
	}
	val := this.kv[key]
	node := this.kn[key]
	this.kl.MoveToBack(node)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if _, ok := this.kv[key]; ok {
		this.kv[key] = value
		node := this.kn[key]
		this.kl.MoveToBack(node)
	} else {
		if this.kl.Len() == this.capacity {
			front := this.kl.Front()
			delete(this.kv, front.Value.(int))
			delete(this.kn, front.Value.(int))
			this.kl.Remove(this.kl.Front())
		}
		this.kv[key] = value
		e := this.kl.PushBack(key)
		this.kn[key] = e
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
