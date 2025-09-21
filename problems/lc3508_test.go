package problems

import (
	"reflect"
	"testing"
)

type RouterOp struct {
	cmd    string
	params []int
	added  bool
	packet []int
	cnt    int
}

func TestRouter(t *testing.T) {
	tests := []struct {
		name        string
		memoryLimit int
		ops         []RouterOp
	}{
		{
			name:        "Example 1",
			memoryLimit: 3,
			ops: []RouterOp{
				{cmd: "addPacket", params: []int{1, 4, 90}, added: true},
				{cmd: "addPacket", params: []int{2, 5, 90}, added: true},
				{cmd: "addPacket", params: []int{1, 4, 90}, added: false},
				{cmd: "addPacket", params: []int{3, 5, 95}, added: true},
				{cmd: "addPacket", params: []int{4, 5, 105}, added: true},
				{cmd: "forwardPacket", packet: []int{2, 5, 90}},
				{cmd: "addPacket", params: []int{5, 2, 110}, added: true},
				{cmd: "getCount", params: []int{5, 100, 110}, cnt: 1},
			},
		},
		{
			name:        "Example 2",
			memoryLimit: 3,
			ops: []RouterOp{
				{cmd: "addPacket", params: []int{1, 4, 90}, added: true},
				{cmd: "getCount", params: []int{5, 100, 110}, cnt: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := RouterConstructor(tt.memoryLimit)
			for _, op := range tt.ops {
				switch op.cmd {
				case "addPacket":
					added := rt.AddPacket(op.params[0], op.params[1], op.params[2])
					if added != op.added {
						t.Errorf("AddPacket(%v) = %v, want %v", op.params, added, op.added)
					}
				case "forwardPacket":
					packet := rt.ForwardPacket()
					if !reflect.DeepEqual(packet, op.packet) {
						t.Errorf("ForwardPacket() = %v, want %v", packet, op.packet)
					}
				case "getCount":
					cnt := rt.GetCount(op.params[0], op.params[1], op.params[2])
					if cnt != op.cnt {
						t.Errorf("GetCount(%v) = %v, want %v", op.params, cnt, op.cnt)
					}
				}
			}
		})
	}
}
