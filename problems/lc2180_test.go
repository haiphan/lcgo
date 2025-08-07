package problems

import (
	"reflect"
	"testing"
)

func TestCountEven(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "Example 1",
			num:  4,
			want: 2,
		},
		{
			name: "Example 2",
			num:  30,
			want: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countEven(tt.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
