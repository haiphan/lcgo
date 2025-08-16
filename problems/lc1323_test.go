package problems

import (
	"reflect"
	"testing"
)

func TestMaximum69Number(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "Example 1",
			num:  9696,
			want: 9996,
		},
		{
			name: "Example 2",
			num:  6,
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum69Number(tt.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
