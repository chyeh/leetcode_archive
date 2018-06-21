package solution

import (
	"container/list"
)

type LRUCache struct {
	dm       map[int]*list.Element //Data Map: Key to Element
	ol       *list.List            // List of used order
	capacity int
}

type kv struct {
	k int
	v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		dm:       make(map[int]*list.Element, capacity),
		ol:       list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.dm[key]
	if !ok {
		return -1
	}
	this.ol.MoveToBack(v)
	return v.Value.(kv).v
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if v, ok := this.dm[key]; ok { // existent: update
		v.Value = kv{key, value}
		this.ol.MoveToBack(v)
	} else { // nonexistent
		if len(this.dm) == this.capacity { // full, invalidate the oldest
			delete(this.dm, this.ol.Front().Value.(kv).k)
			this.ol.Remove(this.ol.Front())
		}
		// insert
		this.ol.PushBack(kv{key, value})
		this.dm[key] = this.ol.Back()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
