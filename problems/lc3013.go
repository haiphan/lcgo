package problems

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type minTracker struct {
	n, k       int
	currentSum int
	count      int // Global counter of items added

	// leftSet: Smallest K items
	leftSet   *redblacktree.Tree
	leftCount int
	leftFreq  map[int]int

	// rightSet: Remaining items
	rightSet  *redblacktree.Tree
	rightFreq map[int]int

	history []int // Circular buffer for eviction
}

func NewMinTracker(n, k int) *minTracker {
	return &minTracker{
		n:         n,
		k:         k,
		leftSet:   redblacktree.NewWithIntComparator(),
		rightSet:  redblacktree.NewWithIntComparator(),
		leftFreq:  make(map[int]int),
		rightFreq: make(map[int]int),
		history:   make([]int, n),
	}
}

func (m *minTracker) Sum() int {
	return m.currentSum
}

func (m *minTracker) Add(x int) {
	// 1. Evict oldest item if window is full
	if m.count >= m.n {
		old := m.history[m.count%m.n]
		if m.leftFreq[old] > 0 {
			m.removeFromLeft(old)
		} else {
			m.removeFromRight(old)
		}
	}

	// 2. Add new item (tentatively to the right)
	m.history[m.count%m.n] = x
	m.addToRight(x)
	m.count++

	// 3. Balance the sets
	m.balance()
}

// --- Internal Helper Methods ---

func (m *minTracker) balance() {
	// Pull into left if it's not full
	for m.leftCount < m.k && m.rightSet.Size() > 0 {
		minVal := m.rightSet.Left().Key.(int)
		m.removeFromRight(minVal)
		m.addToLeft(minVal)
	}

	// Swap if smallest in right is better than largest in left
	if m.leftCount == m.k && m.rightSet.Size() > 0 {
		lMax := m.leftSet.Right().Key.(int)
		rMin := m.rightSet.Left().Key.(int)

		if rMin < lMax {
			m.removeFromLeft(lMax)
			m.removeFromRight(rMin)
			m.addToLeft(rMin)
			m.addToRight(lMax)
		}
	}
}

func (m *minTracker) addToLeft(x int) {
	m.leftFreq[x]++
	if m.leftFreq[x] == 1 {
		m.leftSet.Put(x, nil)
	}
	m.leftCount++
	m.currentSum += x
}

func (m *minTracker) removeFromLeft(x int) {
	m.leftFreq[x]--
	if m.leftFreq[x] == 0 {
		m.leftSet.Remove(x)
	}
	m.leftCount--
	m.currentSum -= x
}

func (m *minTracker) addToRight(x int) {
	m.rightFreq[x]++
	if m.rightFreq[x] == 1 {
		m.rightSet.Put(x, nil)
	}
}

func (m *minTracker) removeFromRight(x int) {
	m.rightFreq[x]--
	if m.rightFreq[x] == 0 {
		m.rightSet.Remove(x)
	}
}

func minimumCost3013(nums []int, k int, dist int) int64 {
	n := len(nums)
	if k == n {
		ans := 0
		for _, x := range nums {
			ans += x
		}
		return int64(ans)
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
