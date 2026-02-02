package problems

import "math"

const (
	NONEHEAP  = 0
	LEFTSIDE  = 1 // Smallest K (Max-Heap)
	RIGHTSIDE = 2 // The Rest (Min-Heap)
)

type ValueItem struct {
	value     int
	heapIndex int
	side      int
}

type minTracker struct {
	n, k       int
	currentSum int
	count      int
	leftHeap   []*ValueItem
	rightHeap  []*ValueItem
	history    []*ValueItem
}

func NewMinTracker(n, k int) *minTracker {
	return &minTracker{
		n:       n,
		k:       k,
		history: make([]*ValueItem, n),
	}
}

func (m *minTracker) Sum() int {
	return m.currentSum
}

func (m *minTracker) Add(x int) {
	pos := m.count % m.n

	// 1. Remove oldest item from its current heap
	if m.count >= m.n {
		old := m.history[pos]
		if old.side == LEFTSIDE {
			m.removeFromLeft(old.heapIndex)
		} else {
			m.removeFromRight(old.heapIndex)
		}
	}

	// 2. Add new item
	newItem := &ValueItem{value: x}
	m.addToRight(newItem)
	m.history[pos] = newItem
	m.count++

	// 3. Balance: Ensure Left has K smallest
	m.balance()
}

func (m *minTracker) balance() {
	// Fill left if it has fewer than K items
	for len(m.leftHeap) < m.k && len(m.rightHeap) > 0 {
		m.addToLeft(m.popMinRight())
	}

	// If Left is full, swap if Right's min is smaller than Left's max
	if len(m.leftHeap) == m.k && len(m.rightHeap) > 0 {
		if m.rightHeap[0].value < m.leftHeap[0].value {
			lMax := m.popMaxLeft()
			rMin := m.popMinRight()
			m.addToLeft(rMin)
			m.addToRight(lMax)
		}
	}
}

// --- Heap Management ---

func (m *minTracker) addToLeft(item *ValueItem) {
	item.side = LEFTSIDE
	item.heapIndex = len(m.leftHeap)
	m.leftHeap = append(m.leftHeap, item)
	m.upHeapLeft(len(m.leftHeap) - 1)
	m.currentSum += item.value
}

func (m *minTracker) addToRight(item *ValueItem) {
	item.side = RIGHTSIDE
	item.heapIndex = len(m.rightHeap)
	m.rightHeap = append(m.rightHeap, item)
	m.upHeapRight(len(m.rightHeap) - 1)
}

func (m *minTracker) removeFromLeft(idx int) {
	m.currentSum -= m.leftHeap[idx].value
	last := len(m.leftHeap) - 1
	if idx != last {
		m.swapLeft(idx, last)
		m.leftHeap = m.leftHeap[:last]
		m.upHeapLeft(idx)
		m.downHeapLeft(idx)
	} else {
		m.leftHeap = m.leftHeap[:last]
	}
}

func (m *minTracker) removeFromRight(idx int) {
	last := len(m.rightHeap) - 1
	if idx != last {
		m.swapRight(idx, last)
		m.rightHeap = m.rightHeap[:last]
		m.upHeapRight(idx)
		m.downHeapRight(idx)
	} else {
		m.rightHeap = m.rightHeap[:last]
	}
}

func (m *minTracker) popMaxLeft() *ValueItem {
	item := m.leftHeap[0]
	m.removeFromLeft(0)
	return item
}

func (m *minTracker) popMinRight() *ValueItem {
	item := m.rightHeap[0]
	m.removeFromRight(0)
	return item
}

// Max-Heap Utilities
func (m *minTracker) upHeapLeft(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if m.leftHeap[i].value <= m.leftHeap[p].value {
			break
		}
		m.swapLeft(i, p)
		i = p
	}
}

func (m *minTracker) downHeapLeft(i int) {
	for {
		l, r := 2*i+1, 2*i+2
		largest := i
		if l < len(m.leftHeap) && m.leftHeap[l].value > m.leftHeap[largest].value {
			largest = l
		}
		if r < len(m.leftHeap) && m.leftHeap[r].value > m.leftHeap[largest].value {
			largest = r
		}
		if largest == i {
			break
		}
		m.swapLeft(i, largest)
		i = largest
	}
}

// Min-Heap Utilities
func (m *minTracker) upHeapRight(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if m.rightHeap[i].value >= m.rightHeap[p].value {
			break
		}
		m.swapRight(i, p)
		i = p
	}
}

func (m *minTracker) downHeapRight(i int) {
	for {
		l, r := 2*i+1, 2*i+2
		smallest := i
		if l < len(m.rightHeap) && m.rightHeap[l].value < m.rightHeap[smallest].value {
			smallest = l
		}
		if r < len(m.rightHeap) && m.rightHeap[r].value < m.rightHeap[smallest].value {
			smallest = r
		}
		if smallest == i {
			break
		}
		m.swapRight(i, smallest)
		i = smallest
	}
}

func (m *minTracker) swapLeft(i, j int) {
	m.leftHeap[i], m.leftHeap[j] = m.leftHeap[j], m.leftHeap[i]
	m.leftHeap[i].heapIndex, m.leftHeap[j].heapIndex = i, j
}

func (m *minTracker) swapRight(i, j int) {
	m.rightHeap[i], m.rightHeap[j] = m.rightHeap[j], m.rightHeap[i]
	m.rightHeap[i].heapIndex, m.rightHeap[j].heapIndex = i, j
}

func minimumCost3013(nums []int, k int, dist int) int64 {
	n := len(nums)
	k2 := k - 2
	if k2 == dist {
		ans := 0
		for i := 1; i < k; i++ {
			ans += nums[i]
		}
		curSum := ans
		for i := 2; i+k2 < n; i++ {
			curSum += (nums[i+k2] - nums[i-1])
			if curSum < ans {
				ans = curSum
			}
		}
		return int64(ans + nums[0])
	}
	kk := k - 1
	tracker := NewMinTracker(dist+1, kk)
	ans := math.MaxInt
	for i := 1; i < n; i++ {
		tracker.Add(nums[i])
		if tracker.count >= kk {
			ans = min(ans, tracker.Sum())
		}
	}
	return int64(ans + nums[0])
}
