package problems

import (
	"reflect"
	"testing"
)

func TestMinDeletionSize960(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int
	}{
		{
			name: "Example 1",
			strs: []string{"ca", "bb", "ac"},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDeletionSize960(tt.strs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minDeletionSize960() = %v, want %v", got, tt.want)
			}
		})
	}
}
