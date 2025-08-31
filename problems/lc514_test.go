package problems

import (
	"reflect"
	"testing"
)

func TestFindRotateSteps(t *testing.T) {
	tests := []struct {
		name string
		ring string
		key  string
		want int
	}{
		{
			name: "Example 1",
			ring: "godding",
			key:  "gd",
			want: 4,
		},
		{
			name: "Example 2",
			ring: "godding",
			key:  "godding",
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRotateSteps(tt.ring, tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRotateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
