package problems

import "sort"

type YEvent struct {
	y      int
	x1, x2 int
	typ    int
}

type IntervalData struct {
	yStart  int
	width   int
	cumArea int64
}

func separateSquares3454(squares [][]int) float64 {
	sqs := squares
	n := len(sqs)
	if n == 0 {
		return 0
	}

	events := make([]YEvent, 0, 2*n)
	xCoordsMap := make(map[int]bool)
	for _, s := range sqs {
		x, y, l := s[0], s[1], s[2]
		events = append(events, YEvent{y, x, x + l, 1})
		events = append(events, YEvent{y + l, x, x + l, -1})
		xCoordsMap[x] = true
		xCoordsMap[x+l] = true
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	sortedX := make([]int, 0, len(xCoordsMap))
	for x := range xCoordsMap {
		sortedX = append(sortedX, x)
	}
	sort.Ints(sortedX)

	xMap := make(map[int]int)
	for i, x := range sortedX {
		xMap[x] = i
	}

	numX := len(sortedX)
	count := make([]int, 4*numX)
	widthTotal := make([]int, 4*numX)

	// Segment Tree Update Function
	update := func(lIdx, rIdx, val int) {
		var updateTree func(int, int, int)
		updateTree = func(node, start, end int) {
			if lIdx <= start && end <= rIdx {
				count[node] += val
			} else {
				mid := (start + end) / 2
				if lIdx < mid {
					updateTree(2*node, start, mid)
				}
				if rIdx > mid {
					updateTree(2*node+1, mid, end)
				}
			}
			if count[node] > 0 {
				widthTotal[node] = sortedX[end] - sortedX[start]
			} else if end-start > 1 {
				widthTotal[node] = widthTotal[2*node] + widthTotal[2*node+1]
			} else {
				widthTotal[node] = 0
			}
		}
		updateTree(1, 0, numX-1)
	}

	// First Pass: Sweep and Cache
	history := make([]IntervalData, 0, len(events))
	var currentArea int64

	for i := 0; i < len(events)-1; i++ {
		update(xMap[events[i].x1], xMap[events[i].x2], events[i].typ)

		w := widthTotal[1]
		dy := int64(events[i+1].y - events[i].y)

		// We only care about intervals with actual height
		if dy > 0 {
			currentArea += int64(w) * dy
			history = append(history, IntervalData{
				yStart:  events[i].y,
				width:   w,
				cumArea: currentArea,
			})
		}
	}

	if len(history) == 0 {
		return 0
	}

	totalArea := history[len(history)-1].cumArea
	target := float64(totalArea) / 2.0

	// Binary Search the cached history
	idx := sort.Search(len(history), func(i int) bool {
		return float64(history[i].cumArea) >= target
	})

	// Linear Interpolation in the identified interval
	found := history[idx]
	areaBefore := float64(0)
	if idx > 0 {
		areaBefore = float64(history[idx-1].cumArea)
	}

	neededArea := target - areaBefore
	return float64(found.yStart) + (neededArea / float64(found.width))
}
