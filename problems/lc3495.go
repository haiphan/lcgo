package problems

import "math/bits"

func minOperationsQueries(queries [][]int) int64 {
	ans := 0
	for _, q := range queries {
		l, r := q[0], q[1]
		pow := (bits.Len64(uint64(l)) - 1) / 2
		l0 := 1 << (2 * pow)
		step := pow + 1
		cnt := 0
		for ; l0 <= r; l0, step = 4*l0, step+1 {
			r0 := (l0 << 2) - 1
			if r0 >= l {
				start := max(l0, l)
				end := min(r0, r)
				cnt += (end - start + 1) * step
			}
		}
		ans += (cnt >> 1) + (cnt & 1)
	}
	return int64(ans)
}
