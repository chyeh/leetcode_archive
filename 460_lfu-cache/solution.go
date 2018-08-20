package solution

import (
	"container/list"
)

type LFUCache struct {
	fn       map[int]*list.Element
	kf       map[int]int
	kn       map[int]*list.Element
	kv       map[int]int
	kl       *list.List
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		fn:       make(map[int]*list.Element),
		kf:       make(map[int]int),
		kn:       make(map[int]*list.Element),
		kv:       make(map[int]int),
		kl:       list.New(),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if _, ok := this.kv[key]; !ok {
		return -1
	}
	this.updateFrequency(key)
	return this.kv[key]
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if _, ok := this.kv[key]; ok {
		this.kv[key] = value
		this.updateFrequency(key)
	} else {
		if this.kl.Len() == this.capacity {
			node := this.kl.Front()
			if this.fn[this.kf[node.Value.(int)]] == node {
				delete(this.fn, this.kf[node.Value.(int)])
			}
			delete(this.kf, node.Value.(int))
			delete(this.kn, node.Value.(int))
			delete(this.kv, node.Value.(int))
			this.kl.Remove(node)
		}
		this.kv[key] = value
		e := this.kl.PushFront(key)
		this.fn[0] = e
		this.kn[key] = e
		this.kf[key] = 0
		this.updateFrequency(key)
	}
}

func (this *LFUCache) updateFrequency(key int) {
	node := this.kn[key]
	if this.fn[this.kf[key]] == node {
		if node.Prev() != nil && this.kf[node.Prev().Value.(int)] == this.kf[key] {
			this.fn[this.kf[key]] = node.Prev()
		} else {
			delete(this.fn, this.kf[key])
		}
	}
	this.kf[key]++
	if _, ok := this.fn[this.kf[key]]; ok {
		this.kl.MoveAfter(node, this.fn[this.kf[key]])
		this.fn[this.kf[key]] = node
	} else if _, ok := this.fn[this.kf[key]-1]; ok {
		this.kl.MoveAfter(node, this.fn[this.kf[key]-1])
	}
	this.fn[this.kf[key]] = node
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
