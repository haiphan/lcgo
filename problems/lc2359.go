package problems

func explore(cur int, edges []int, dist []int) {
	d := 0
	for {
		if cur == -1 || dist[cur] <= d {
			break
		}
		dist[cur] = d
		d++
		cur = edges[cur]
	}
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	N := len(edges)
	dist1, dist2 := make([]int, N), make([]int, N)
	for i := range N {
		dist1[i], dist2[i] = N, N
	}
	explore(node1, edges, dist1)
	explore(node2, edges, dist2)
	res, md := -1, N
	for i := range N {
		cd := max(dist1[i], dist2[i])
		if cd < md {
			md = cd
			res = i
		}
	}
	return res
}
