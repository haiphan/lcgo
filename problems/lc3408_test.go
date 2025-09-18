package problems

import (
	"testing"
)

type TMOp struct {
	op     string
	params []int
	want   int
}

func TestTaskManager(t *testing.T) {
	tests := []struct {
		name  string
		tasks [][]int
		ops   []TMOp
	}{
		{
			name:  "Example 1",
			tasks: [][]int{{1, 101, 8}, {2, 102, 20}, {3, 103, 5}},
			ops: []TMOp{
				{op: "add", params: []int{4, 104, 5}},
				{op: "edit", params: []int{102, 9}},
				{op: "execTop", want: 2},
				{op: "rmv", params: []int{101}},
				{op: "add", params: []int{50, 101, 8}},
				{op: "execTop", want: 50},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := TaskManagerConstructor(tt.tasks)
			for _, op := range tt.ops {
				switch op.op {
				case "add":
					tm.Add(op.params[0], op.params[1], op.params[2])
				case "edit":
					tm.Edit(op.params[0], op.params[1])
				case "execTop":
					got := tm.ExecTop()
					if got != op.want {
						t.Errorf("TaskManager() = %v, want %v", got, op.want)
					}
				case "rmv":
					tm.Rmv(op.params[0])
				}
			}
		})
	}
}
