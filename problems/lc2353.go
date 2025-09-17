package problems

type Food struct {
	n, c string
	r    int
}

type FoodRatings struct {
	cToQ map[string][]Food
	nToF map[string]Food
}

func FLessThan(a, b Food) bool {
	if a.r == b.r {
		return a.n < b.n
	}
	return b.r < a.r
}

func FHPush(h []Food, x Food) []Food {
	h = append(h, x)
	cur := len(h) - 1
	for cur > 0 {
		p := (cur - 1) >> 1
		if FLessThan(h[p], h[cur]) {
			break
		}
		h[p], h[cur] = h[cur], h[p]
		cur = p
	}
	return h
}

func FHPop(h []Food) []Food {
	last := h[len(h)-1]
	h[0] = last
	h = h[:len(h)-1]
	cur, l := 0, 1
	for l < len(h) {
		r := l + 1
		c := l
		if r < len(h) && FLessThan(h[r], h[l]) {
			c = r
		}
		if FLessThan(h[cur], h[c]) {
			break
		}
		h[cur], h[c] = h[c], h[cur]
		cur = c
		l = cur*2 + 1
	}
	return h
}

func FoodRatingsConstructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	cToQ := make(map[string][]Food, len(foods))
	nToF := make(map[string]Food, len(foods))
	for i, f := range foods {
		item := Food{n: f, c: cuisines[i], r: ratings[i]}
		if _, has := cToQ[cuisines[i]]; !has {
			cToQ[cuisines[i]] = make([]Food, 0)
		}
		cToQ[cuisines[i]] = FHPush(cToQ[cuisines[i]], Food{n: f, c: cuisines[i], r: ratings[i]})
		nToF[f] = item
	}
	return FoodRatings{cToQ: cToQ, nToF: nToF}
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	item := fr.nToF[food]
	item.r = newRating
	fr.nToF[food] = item
	copy := item
	fr.cToQ[copy.c] = FHPush(fr.cToQ[copy.c], copy)
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	top := fr.cToQ[cuisine][0]
	for top.r != fr.nToF[top.n].r {
		fr.cToQ[cuisine] = FHPop(fr.cToQ[cuisine])
		top = fr.cToQ[cuisine][0]
	}
	return top.n
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
