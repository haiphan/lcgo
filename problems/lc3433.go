package problems

import (
	"sort"
	"strconv"
	"strings"
)

type MEV struct {
	t int
	y int // 0 OFFLINE
	c string
}

func countMentions(numberOfUsers int, events [][]string) []int {
	E, N := len(events), numberOfUsers
	res := make([]int, N)
	mevs := make([]MEV, E)
	online := make([]int, N)
	for i, e := range events {
		y := 0
		if e[0] == "MESSAGE" {
			y = 1
		}
		t, _ := strconv.Atoi(e[1])
		mevs[i] = MEV{t, y, e[2]}
	}
	sort.Slice(mevs, func(i, j int) bool {
		if mevs[i].t == mevs[j].t {
			return mevs[i].y < mevs[j].y
		}
		return mevs[i].t < mevs[j].t
	})
	for _, e := range mevs {
		if e.y == 0 {
			uid, _ := strconv.Atoi(e.c)
			online[uid] = e.t + 60
		} else {
			switch e.c {
			case "ALL":
				for i := range N {
					res[i]++
				}
			case "HERE":
				for i := range N {
					if online[i] <= e.t {
						res[i]++
					}
				}
			default:
				arr := strings.Split(e.c, " ")
				for _, idstr := range arr {
					uid, _ := strconv.Atoi(idstr[2:])
					res[uid]++
				}
			}
		}
	}
	return res
}
