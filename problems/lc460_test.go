package problems

import (
	"testing"
)

type LCOperation struct {
	Op    string
	Key   int
	Value int
	want  int
}

func TestLFUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []LCOperation
	}{
		{
			name:     "Example 1",
			capacity: 2,
			ops: []LCOperation{
				{Op: "put", Key: 1, Value: 1},
				{Op: "put", Key: 2, Value: 2},
				{Op: "get", Key: 1, want: 1},
				{Op: "put", Key: 3, Value: 3},
				{Op: "get", Key: 2, want: -1},
				{Op: "get", Key: 3, want: 3},
				{Op: "put", Key: 4, Value: 4},
				{Op: "get", Key: 1, want: -1},
				{Op: "get", Key: 3, want: 3},
				{Op: "get", Key: 4, want: 4},
			},
		},
		{
			name:     "Example 2",
			capacity: 2,
			ops: []LCOperation{
				{Op: "get", Key: 2, want: -1},
				{Op: "put", Key: 2, Value: 6},
				{Op: "get", Key: 1, want: -1},
				{Op: "put", Key: 1, Value: 5},
				{Op: "put", Key: 1, Value: 2},
				{Op: "get", Key: 1, want: 2},
				{Op: "get", Key: 2, want: 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := LFUCacheConstructor(tt.capacity)
			for _, op := range tt.ops {
				switch op.Op {
				case "put":
					lc.Put(op.Key, op.Value)
				case "get":
					got := lc.Get(op.Key)
					if got != op.want {
						t.Errorf("Get(%d) = %d, want %d", op.Key, got, op.want)
					}
				}
			}
		})
	}
}
