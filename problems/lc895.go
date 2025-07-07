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

func (fstack *FreqStack) Push(val int) {
	fstack.ntof[val]++
	fstack.mf = max(fstack.mf, fstack.ntof[val])
	if fstack.ntof[val] == len(fstack.stack) {
		fstack.stack = append(fstack.stack, []int{})
	}
	fstack.stack[fstack.ntof[val]] = append(fstack.stack[fstack.ntof[val]], val)
}

func (fstack *FreqStack) Pop() int {
	v := fstack.stack[fstack.mf][len(fstack.stack[fstack.mf])-1]
	fstack.ntof[v]--
	fstack.stack[fstack.mf] = fstack.stack[fstack.mf][:len(fstack.stack[fstack.mf])-1]
	if len(fstack.stack[fstack.mf]) == 0 {
		fstack.mf--
	}
	return v
}
