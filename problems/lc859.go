package problems

type FreqStack struct {
	mf    int
	ntof  map[int]int
	stack [][]int
}

func Constructor() FreqStack {
	stack := make([][]int, 1)
	stack[0] = []int{}
	return FreqStack{mf: 0, ntof: make(map[int]int), stack: stack}
}

func (this *FreqStack) Push(val int) {
	this.ntof[val]++
	this.mf = max(this.mf, this.ntof[val])
	if this.ntof[val] == len(this.stack) {
		this.stack = append(this.stack, []int{})
	}
	this.stack[this.ntof[val]] = append(this.stack[this.ntof[val]], val)
}

func (this *FreqStack) Pop() int {
	v := this.stack[this.mf][len(this.stack[this.mf])-1]
	this.ntof[v]--
	this.stack[this.mf] = this.stack[this.mf][:len(this.stack[this.mf])-1]
	if len(this.stack[this.mf]) == 0 {
		this.mf--
	}
	return v
}
