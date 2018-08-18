package solution

type node struct {
	val  int
	next *node
	prev *node
}

type MyLinkedList struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	n := &node{}
	n.prev = n
	n.next = n
	return MyLinkedList{
		root: n,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	i, iterator := 0, this.root.next
	for ; i < index && iterator != this.root; i, iterator = i+1, iterator.next {
	}
	if i == index && iterator != this.root {
		return iterator.val
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	n := &node{
		val:  val,
		prev: this.root,
		next: this.root.next,
	}
	this.root.next.prev = n
	this.root.next = n
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	n := &node{
		val:  val,
		prev: this.root.prev,
		next: this.root,
	}
	this.root.prev.next = n
	this.root.prev = n
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	i, iterator := 0, this.root.next
	for ; i < index && iterator != this.root; i, iterator = i+1, iterator.next {
	}
	if i == index && iterator == this.root {
		this.AddAtTail(val)
	} else if i == index {
		n := &node{
			val:  val,
			prev: iterator.prev,
			next: iterator,
		}
		iterator.prev.next = n
		iterator.prev = n
	}
	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	i, iterator := 0, this.root.next
	for ; i < index && iterator != this.root; i, iterator = i+1, iterator.next {
	}
	if i == index && iterator != this.root {
		iterator.prev.next = iterator.next
		iterator.next.prev = iterator.prev
	}
	return
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
