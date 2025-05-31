package problems

func snakesAndLadders(board [][]int) int {
	N := len(board)
	MAXV := N * N
	GOAL := MAXV - 1
	fb := make([]int, MAXV)
	i := 0
	for r := N - 1; r >= 0; r-- {
		odd := (N-1-r)%2 == 1
		for c := 0; c < N; c++ {
			rc := c
			if odd {
				rc = N - 1 - c
			}
			fb[i] = board[r][rc]
			i++
		}
	}
	dist := make([]int, MAXV)
	q := make([]int, MAXV)
	dist[0] = 0
	l, r := 0, 1
	for i := 1; i < MAXV; i++ {
		dist[i] = MAXV
		q[i] = -1
	}
	for l < r {
		cur := q[l]
		l++
		for i := 1; i <= 6; i++ {
			next := cur + i
			if next > GOAL {
				break
			}
			if fb[next] != -1 {
				next = fb[next] - 1
			}
			if next == GOAL {
				return dist[cur] + 1
			}
			if dist[next] > dist[cur]+1 {
				dist[next] = dist[cur] + 1
				q[r] = next
				r++
			}
		}
	}
	return -1
}
