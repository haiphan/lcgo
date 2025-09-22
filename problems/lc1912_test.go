package problems

import (
	"reflect"
	"testing"
)

type MRSOp struct {
	cmd    string
	params []int
	shops  []int
	rep    [][]int
}

func TestMovieRentingSystem(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		entries [][]int
		ops     []MRSOp
	}{
		{
			name:    "Example 1",
			n:       3,
			entries: [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}},
			ops: []MRSOp{
				{cmd: "search", params: []int{1}, shops: []int{1, 0, 2}},
				{cmd: "rent", params: []int{0, 1}},
				{cmd: "rent", params: []int{1, 2}},
				{cmd: "report", rep: [][]int{{0, 1}, {1, 2}}},
				{cmd: "drop", params: []int{1, 2}},
				{cmd: "search", params: []int{2}, shops: []int{0, 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mrs := MovieRentingSystemConstructor(tt.n, tt.entries)
			for _, op := range tt.ops {
				switch op.cmd {
				case "search":
					shops := mrs.Search(op.params[0])
					if !reflect.DeepEqual(shops, op.shops) {
						t.Errorf("Search(%v) = %v, want %v", op.params, shops, op.shops)
					}
				case "rent":
					mrs.Rent(op.params[0], op.params[1])
				case "drop":
					mrs.Drop(op.params[0], op.params[1])
				case "report":
					rep := mrs.Report()
					if !reflect.DeepEqual(rep, op.rep) {
						t.Errorf("Report() = %v, want %v", rep, op.rep)
					}
				}
			}
		})
	}
}
