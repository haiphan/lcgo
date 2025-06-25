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

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.stack) > 0 && getPrice(this.stack[len(this.stack)-1]) <= price {
		span += getSpan(this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, span*SBASE+price)
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := StockSpannerConstructor();
 * param_1 := obj.Next(price);
 */
