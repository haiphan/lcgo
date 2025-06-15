package problems

import (
	"reflect"
	"testing"
)

func TestMaxDiff(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "Example 1",
			num:  555,
			want: 888,
		},
		{
			name: "Example 2",
			num:  9,
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDiff(tt.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
