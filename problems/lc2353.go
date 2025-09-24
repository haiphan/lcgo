package problems

import "container/heap"

type Food struct {
	n, c         string
	r, heapIndex int
}

type FHeap []*Food

func (fh FHeap) Len() int { return len(fh) }
func (fh FHeap) Less(i, j int) bool {
	return fh[i].r > fh[j].r || (fh[i].r == fh[j].r && fh[i].n < fh[j].n)
}
func (fh FHeap) Swap(i, j int) {
	fh[i], fh[j] = fh[j], fh[i]
	fh[i].heapIndex = i
	fh[j].heapIndex = j
}

func (fh *FHeap) Push(x interface{}) {
	item := x.(*Food)
	*fh = append(*fh, item)
}

func (fh *FHeap) Pop() interface{} {
	old := *fh
	n := len(old)
	item := old[n-1]
	*fh = old[0 : n-1]
	return item
}

type FoodRatings struct {
	cToQ map[string]*FHeap
	nToF map[string]*Food
}

func FoodRatingsConstructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	cToQ := make(map[string]*FHeap, len(foods))
	nToF := make(map[string]*Food, len(foods))
	for i, f := range foods {
		item := &Food{n: f, c: cuisines[i], r: ratings[i]}
		if _, has := cToQ[cuisines[i]]; !has {
			cToQ[cuisines[i]] = &FHeap{}
		}
		item.heapIndex = cToQ[cuisines[i]].Len()
		*cToQ[cuisines[i]] = append(*cToQ[cuisines[i]], item)
		nToF[f] = item
	}
	for _, fh := range cToQ {
		heap.Init(fh)
	}
	return FoodRatings{cToQ: cToQ, nToF: nToF}
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	f := fr.nToF[food]
	f.r = newRating
	heap.Fix(fr.cToQ[f.c], f.heapIndex)
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	fh := fr.cToQ[cuisine]
	if fh != nil && fh.Len() > 0 {
		return (*fh)[0].n
	}
	return ""
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
