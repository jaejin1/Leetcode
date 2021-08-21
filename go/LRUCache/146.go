package LRUCache

type LRUCache struct {
	capacity   int
	cache      map[int]*node
	head, tail *node
}

type node struct {
	prev, next *node
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    map[int]*node{},
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]

	if !ok {
		return -1
	}

	// cache hit on most recently used; do nothing
	if this.head != nil && node.key == this.head.key {
		return node.value
	}
	// cache hit on least recently used; pop to front
	if this.tail != nil && node.key == this.tail.key {
		this.pop()
		this.insert(node)
		this.cache[node.key] = node
		return node.value
	}
	// pop to front
	if node.prev != nil && node.next != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
		this.insert(node)
	}

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.cache[key].value = value
		return
	}

	n := &node{key: key, value: value}
	this.cache[key] = n
	this.insert(n)

	if len(this.cache) > this.capacity {
		this.pop()
	}

	return

}

func (this *LRUCache) insert(n *node) {
	// nothing inserted yet
	if this.tail == nil {
		this.tail = n
		this.head = n
		n.next = this.tail
		return
	}
	this.head.prev = n
	n.next = this.head
	this.head = n
}

func (this *LRUCache) pop() {
	// remove key
	delete(this.cache, this.tail.key)
	// nothing in cache yet
	if this.tail.prev == nil {
		this.tail = nil
		return
	}
	// pop tail
	this.tail.prev.next = nil
	this.tail = this.tail.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
