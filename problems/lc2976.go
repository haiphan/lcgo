package problems

func minimumCost2976(source string, target string, original []byte, changed []byte, cost []int) int64 {
	BIG := 4000000
	var dist [26][26]int
	for i := range 26 {
		for j := range 26 {
			if i == j {
				continue
			}
			dist[i][j] = BIG
		}
	}
	for i, c := range original {
		si, di := c-'a', changed[i]-'a'
		dist[si][di] = min(dist[si][di], cost[i])
	}
	for k := range 26 {
		for i := range 26 {
			for j := range 26 {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	ans := 0
	for i, c := range source {
		d := dist[c-'a'][target[i]-'a']
		if d >= BIG {
			return -1
		}
		ans += d
	}
	return int64(ans)
}
