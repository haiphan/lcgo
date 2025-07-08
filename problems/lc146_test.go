package problems

import (
	"testing"
)

type LRUOp struct {
	op    string
	key   int
	value int
	want  int
}

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []LRUOp
	}{
		{
			name:     "Example 1",
			capacity: 2,
			ops: []LRUOp{
				{op: "put", key: 1, value: 1},
				{op: "put", key: 2, value: 2},
				{op: "get", key: 1, want: 1},
				{op: "put", key: 3, value: 3},
				{op: "get", key: 2, want: -1},
			},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lrucache := LRUCacheConstructor(tt.capacity)
			for _, op := range tt.ops {
				switch op.op {
				case "put":
					lrucache.Put(op.key, op.value)
				case "get":
					got := lrucache.Get(op.key)
					if got != op.want {
						t.Errorf("Get(%d) = %d, want %d", op.key, got, op.want)
					}
				default:
					t.Errorf("Unknown operation: %s", op.op)
				}
			}
		})
	}
}
