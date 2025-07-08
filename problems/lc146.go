package problems

type LNode struct {
	key   int
	value int
	next  *LNode
	prev  *LNode
}

type LRUCache struct {
	cap   int
	left  *LNode
	right *LNode
	cache map[int]*LNode
}

func LRUCacheConstructor(capacity int) LRUCache {
	left := &LNode{key: -1, value: -1}
	right := &LNode{key: -1, value: -1}
	left.next = right
	right.prev = left
	return LRUCache{cap: capacity, cache: make(map[int]*LNode), left: left, right: right}
}

func moveBack(node *LNode, r *LNode) {
	n, p := node.next, node.prev
	if r == n {
		return
	}
	if n != nil {
		p.next = n
		n.prev = p
	}
	newPrev, newNext := r.prev, r
	newPrev.next, newNext.prev = node, node
	node.prev, node.next = newPrev, newNext
}

func (this *LRUCache) Get(key int) int {
	node, has := this.cache[key]
	if !has {
		return -1
	}
	moveBack(node, this.right)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, has := this.cache[key]
	if has {
		node.value = value
	} else {
		node = &LNode{key: key, value: value}
		this.cache[key] = node
	}
	moveBack(node, this.right)
	if len(this.cache) > this.cap {
		node = this.left.next
		next := node.next
		this.left.next = next
		next.prev = this.left
		delete(this.cache, node.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
