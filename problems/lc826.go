package problems

import "sort"

type JobItem struct {
	p int
	d int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	jobs := make([]JobItem, len(difficulty))
	for i, d := range difficulty {
		jobs[i] = JobItem{p: profit[i], d: d}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].d < jobs[j].d
	})
	sort.Ints(worker)
	j, res := 0, 0
	best := 0
	for _, w := range worker {
		for j < len(jobs) && jobs[j].d <= w {
			best = max(best, jobs[j].p)
			j++
		}
		res += best
	}
	return res
}
