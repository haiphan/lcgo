package problems

import "github.com/emirpasic/gods/trees/redblacktree"

const MBASE int = 10001
const SHBASE int = 300000

func PSMToK(p, s, m int) int {
	return m + (MBASE * s) + (MBASE * SHBASE * p)
}

func kToPSM(k int) [3]int {
	m := k % MBASE
	s := (k / MBASE) % SHBASE
	p := k / (MBASE * SHBASE)
	return [3]int{p, s, m}
}

func getFive(tree *redblacktree.Tree) []int {
	res := make([]int, 0)
	var inOrder func(cur *redblacktree.Node)
	inOrder = func(cur *redblacktree.Node) {
		if cur == nil || len(res) >= 5 {
			return
		}
		inOrder(cur.Left)
		if len(res) >= 5 {
			return
		}
		res = append(res, cur.Key.(int))
		inOrder(cur.Right)
	}
	inOrder(tree.Root)
	return res
}

type MovieRentingSystem struct {
	SMToP   map[[2]int]int
	mToFree map[int]*redblacktree.Tree
	rented  *redblacktree.Tree
}

// smp
func MovieRentingSystemConstructor(n int, entries [][]int) MovieRentingSystem {
	SMToP := make(map[[2]int]int)
	mToFree := make(map[int]*redblacktree.Tree)
	rented := redblacktree.NewWithIntComparator()
	for _, e := range entries {
		s, m, p := e[0], e[1], e[2]
		k := PSMToK(p, s, m)
		SMToP[[2]int{s, m}] = p
		free, has := mToFree[m]
		if !has {
			free = redblacktree.NewWithIntComparator()
			mToFree[m] = free
		}
		free.Put(k, nil)
	}
	return MovieRentingSystem{SMToP: SMToP, mToFree: mToFree, rented: rented}
}

func (mrs *MovieRentingSystem) Search(movie int) []int {
	free, has := mrs.mToFree[movie]
	if !has {
		return []int{}
	}
	res := getFive(free)
	for i, v := range res {
		res[i] = kToPSM(v)[1]
	}
	return res
}

func (mrs *MovieRentingSystem) Rent(shop int, movie int) {
	p := mrs.SMToP[[2]int{shop, movie}]
	k := PSMToK(p, shop, movie)
	mrs.rented.Put(k, nil)
	mrs.mToFree[movie].Remove(k)
}

func (mrs *MovieRentingSystem) Drop(shop int, movie int) {
	p := mrs.SMToP[[2]int{shop, movie}]
	k := PSMToK(p, shop, movie)
	mrs.rented.Remove(k)
	mrs.mToFree[movie].Put(k, nil)
}

func (mrs *MovieRentingSystem) Report() [][]int {
	arr := getFive(mrs.rented)
	res := make([][]int, len(arr))
	for i, k := range arr {
		psm := kToPSM(k)
		res[i] = []int{psm[1], psm[2]}
	}
	return res
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */
