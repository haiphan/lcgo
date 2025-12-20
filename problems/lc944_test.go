package problems

import (
	"reflect"
	"testing"
)

func TestMinDeletionSize(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int
	}{
		{
			name: "Example 1",
			strs: []string{"cba", "daf", "ghi"},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDeletionSize(tt.strs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
