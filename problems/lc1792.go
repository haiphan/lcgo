package problems

import "container/heap"

type Score struct {
	good  int
	total int
	inc   float64
}

// ScoreHeap implements heap.Interface for a max-heap of Scores.
type ScoreHeap []Score

func (sh ScoreHeap) Len() int { return len(sh) }

// Less returns true if element `j` is less than element `i`, creating a max-heap.
func (sh ScoreHeap) Less(i, j int) bool {
	return sh[j].inc < sh[i].inc
}

func (sh ScoreHeap) Swap(i, j int) {
	sh[i], sh[j] = sh[j], sh[i]
}

// Push adds an item to the heap. It uses a pointer receiver to modify the slice.
func (sh *ScoreHeap) Push(x any) {
	*sh = append(*sh, x.(Score))
}

// Pop removes the highest-priority item from the heap. It uses a pointer receiver.
func (sh *ScoreHeap) Pop() any {
	old := *sh
	n := len(old)
	x := old[n-1]
	*sh = old[0 : n-1]
	return x
}

func getInc(a int, b int) float64 {
	return float64(b-a) / float64(b*(b+1))
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	n := len(classes)
	scores := make([]Score, n)
	sum := 0.0
	for i := range n {
		good, total := classes[i][0], classes[i][1]
		sum += float64(good) / float64(total)
		scores[i] = Score{good: good, total: total, inc: getInc(good, total)}
	}

	h := (*ScoreHeap)(&scores)
	heap.Init(h)
	for i := 0; i < extraStudents; i++ {
		sum += (*h)[0].inc
		(*h)[0].good++
		(*h)[0].total++
		(*h)[0].inc = getInc((*h)[0].good, (*h)[0].total)
		heap.Fix(h, 0)
	}
	return sum / float64(n)
}
