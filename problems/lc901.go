package problems

const SBASE int = 1e5 + 1

type StockSpanner struct {
	stack []int
}

func getPrice(x int) int {
	return x % SBASE
}

func getSpan(x int) int {
	return x / SBASE
}

func StockSpannerConstructor() StockSpanner {
	return StockSpanner{stack: make([]int, 0, 1e4)}
}

func (sp *StockSpanner) Next(price int) int {
	span := 1
	for len(sp.stack) > 0 && getPrice(sp.stack[len(sp.stack)-1]) <= price {
		span += getSpan(sp.stack[len(sp.stack)-1])
		sp.stack = sp.stack[:len(sp.stack)-1]
	}
	sp.stack = append(sp.stack, span*SBASE+price)
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := StockSpannerConstructor();
 * param_1 := obj.Next(price);
 */
