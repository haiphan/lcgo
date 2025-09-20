package problems

type LFNode struct {
	key   int
	value int
	cnt   int
	prev  *LFNode
	next  *LFNode
}

type NodeList struct {
	head *LFNode
	tail *LFNode
}

func createNodeList() *NodeList {
	head, tail := createLFNode(-1, -1), createLFNode(-1, -1)
	head.next = tail
	tail.prev = head
	return &NodeList{head: head, tail: tail}
}

func (l *NodeList) isEmpty() bool {
	return l.head.next == l.tail
}

func (l *NodeList) insert(node *LFNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *NodeList) remove(node *LFNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *NodeList) removeTail() *LFNode {
	if l.isEmpty() {
		return nil
	}
	node := l.tail.prev
	l.remove(node)
	return node
}

func createLFNode(key, value int) *LFNode {
	return &LFNode{key: key, value: value, cnt: 1}
}

type LFUCache struct {
	valueMap map[int]*LFNode
	minF     int
	capacity int
	listMap  map[int]*NodeList
}

func LFUCacheConstructor(capacity int) LFUCache {
	return LFUCache{capacity: capacity, minF: 1, valueMap: make(map[int]*LFNode), listMap: make(map[int]*NodeList)}
}

func (lc *LFUCache) counter(node *LFNode) {
	cnt := node.cnt
	nextCnt := cnt + 1
	node.cnt = nextCnt
	curList := lc.listMap[cnt]
	nextList, has := lc.listMap[nextCnt]
	curList.remove(node)
	if !has {
		nextList = createNodeList()
		lc.listMap[nextCnt] = nextList
	}
	nextList.insert(node)
	if lc.minF == cnt && curList.isEmpty() {
		lc.minF++
	}
}

func (this *LFUCache) Get(key int) int {
	node, has := this.valueMap[key]
	if !has {
		return -1
	}
	this.counter(node)
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	node, has := this.valueMap[key]
	if has {
		node.value = value
		this.counter(node)
	} else {
		node = createLFNode(key, value)
		if len(this.valueMap) == this.capacity {
			minList := this.listMap[this.minF]
			rmNode := minList.removeTail()
			delete(this.valueMap, rmNode.key)
		}
		oneList, hasL := this.listMap[1]
		if !hasL {
			oneList = createNodeList()
			this.listMap[1] = oneList
		}
		oneList.insert(node)
		this.valueMap[key] = node
		this.minF = 1
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
