package problems

import "github.com/emirpasic/gods/trees/redblacktree"

type CItem struct {
	v int
	f int
}

func ItemComparator(a, b interface{}) int {
	itemA := a.(CItem)
	itemB := b.(CItem)

	if itemA.f < itemB.f {
		return -1
	}
	if itemA.f > itemB.f {
		return 1
	}

	if itemA.v < itemB.v {
		return -1
	}
	if itemA.v > itemB.v {
		return 1
	}

	return 0
}

func findXSum3321(nums []int, k int, x int) []int64 {
	bigTree := redblacktree.NewWith(ItemComparator)
	smallTree := redblacktree.NewWith(ItemComparator)
	bIt := bigTree.Iterator()
	sIt := smallTree.Iterator()
	n := len(nums)
	nc := make(map[int]int)
	sum := 0
	ans := make([]int64, n-k+1)
	balance := func() {
		for bigTree.Size() < x && smallTree.Size() > 0 {
			sIt.Last()
			item := sIt.Key().(CItem)
			sum += item.f * item.v
			smallTree.Remove(item)
			bigTree.Put(item, true)
		}
		for bigTree.Size() > x {
			bIt.First()
			item := bIt.Key().(CItem)
			sum -= item.f * item.v
			bigTree.Remove(item)
			smallTree.Put(item, true)
		}
		if bigTree.Size() == 0 || smallTree.Size() == 0 {
			return
		}
		for {
			bIt.First()
			bigItem := bIt.Key().(CItem)
			sIt.Last()
			smallItem := sIt.Key().(CItem)
			if ItemComparator(smallItem, bigItem) < 0 {
				return
			}
			bigTree.Remove(bigItem)
			smallTree.Remove(smallItem)
			sum += smallItem.f*smallItem.v - bigItem.f*bigItem.v
			smallTree.Put(bigItem, true)
			bigTree.Put(smallItem, true)
		}
	}

	add := func(v int) {
		cur := CItem{v: v, f: nc[v]}
		if _, has := bigTree.Get(cur); has {
			bigTree.Remove(cur)
			sum -= cur.f * cur.v
		} else {
			smallTree.Remove(cur)
		}
		nc[v]++
		cur.f++
		smallTree.Put(cur, true)
		balance()
	}
	remove := func(v int) {
		cur := CItem{v: v, f: nc[v]}
		if _, has := bigTree.Get(cur); has {
			bigTree.Remove(cur)
			sum -= cur.f * cur.v
		} else {
			smallTree.Remove(cur)
		}
		nc[v]--
		cur.f--
		if cur.f > 0 {
			smallTree.Put(cur, true)
		} else {
			delete(nc, v)
		}
		balance()
	}
	for i := range k {
		add(nums[i])
	}
	ans[0] = int64(sum)
	for i := k; i < n; i++ {
		add(nums[i])
		remove(nums[i-k])
		ans[i-k+1] = int64(sum)
	}
	return ans
}
