package problems

import (
	"container/heap"
	"math"
)

type NodeDist struct {
	node int
	dist int
}

type NDQueue []*NodeDist

func (pq NDQueue) Len() int           { return len(pq) }
func (pq NDQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq NDQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *NDQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*NodeDist))
}

func (pq *NDQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func minCost3650(n int, edges [][]int) int {
	adj := make([][][2]int, n)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w * 2})
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	pq := &NDQueue{}
	heap.Init(pq)
	heap.Push(pq, &NodeDist{node: 0, dist: 0})
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*NodeDist)
		u := curr.node
		d := curr.dist
		if d > dist[u] {
			continue
		}

		if u == n-1 {
			return dist[u]
		}

		for _, neighbor := range adj[u] {
			v, w := neighbor[0], neighbor[1]
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				heap.Push(pq, &NodeDist{node: v, dist: dist[v]})
			}
		}
	}
	return -1
}
