package problems

import (
	"testing"
)

type ExamRoomOp struct {
	cmd     string
	p, want int
}

func TestExamRoom(t *testing.T) {
	tests := []struct {
		name string
		n    int
		ops  []ExamRoomOp
	}{
		{
			name: "Example 1",
			n:    10,
			ops: []ExamRoomOp{
				{cmd: "seat", want: 0},
				{cmd: "seat", want: 9},
				{cmd: "seat", want: 4},
				{cmd: "seat", want: 2},
				{cmd: "leave", p: 4},
				{cmd: "seat", want: 5},
			},
		},
		{
			name: "Example 2",
			n:    8,
			ops: []ExamRoomOp{
				{cmd: "seat", want: 0},
				{cmd: "seat", want: 7},
				{cmd: "seat", want: 3},
				{cmd: "leave", p: 0},
				{cmd: "leave", p: 7},
				{cmd: "seat", want: 7},
				{cmd: "seat", want: 0},
				{cmd: "seat", want: 5},
				{cmd: "seat", want: 1},
				{cmd: "seat", want: 2},
				{cmd: "seat", want: 4},
				{cmd: "seat", want: 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eRoom := ExamRoomConstructor(tt.n)
			for _, op := range tt.ops {
				switch op.cmd {
				case "seat":
					got := eRoom.Seat()
					if got != op.want {
						t.Errorf("ExamRoom.Seat() = %v, want %v", got, op.want)
					}
				case "leave":
					eRoom.Leave(op.p)
				}
			}
		})
	}
}
