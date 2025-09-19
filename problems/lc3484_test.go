package problems

import (
	"testing"
)

type SpreadSheetOp struct {
	op      string
	cell    string
	formula string
	value   int
	want    int
}

func TestSpreadsheet(t *testing.T) {
	tests := []struct {
		name string
		rows int
		ops  []SpreadSheetOp
	}{
		{
			name: "Example 1",
			rows: 3,
			ops: []SpreadSheetOp{
				{op: "getValue", formula: "=5+7", want: 12},
				{op: "setCell", cell: "A1", value: 10},
				{op: "getValue", formula: "=A1+6", want: 16},
				{op: "setCell", cell: "B2", value: 15},
				{op: "getValue", formula: "=A1+B2", want: 25},
				{op: "resetCell", cell: "A1"},
				{op: "getValue", formula: "=A1+B2", want: 15},
			},
		},
		{
			name: "Example 2",
			rows: 36,
			ops: []SpreadSheetOp{
				{op: "getValue", formula: "=72260+45645", want: 117905},
				{op: "getValue", formula: "=K36+29561", want: 29561},
				{op: "resetCell", cell: "U13"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := SpreadsheetConstructor(tt.rows)
			for _, op := range tt.ops {
				switch op.op {
				case "setCell":
					ss.SetCell(op.cell, op.value)
				case "getValue":
					if got := ss.GetValue(op.formula); got != op.want {
						t.Errorf("GetValue(%v) = %v, want %v", op.formula, got, op.want)
					}
				case "resetCell":
					ss.ResetCell(op.cell)
				}
			}
		})
	}
}
