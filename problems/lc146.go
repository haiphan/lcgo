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

func (lruc *LRUCache) Get(key int) int {
	node, has := lruc.cache[key]
	if !has {
		return -1
	}
	moveBack(node, lruc.right)
	return node.value
}

func (lruc *LRUCache) Put(key int, value int) {
	node, has := lruc.cache[key]
	if has {
		node.value = value
	} else {
		node = &LNode{key: key, value: value}
		lruc.cache[key] = node
	}
	moveBack(node, lruc.right)
	if len(lruc.cache) > lruc.cap {
		node = lruc.left.next
		next := node.next
		lruc.left.next = next
		next.prev = lruc.left
		delete(lruc.cache, node.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
