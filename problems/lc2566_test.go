package problems

import (
	"reflect"
	"testing"
)

func TestMinMaxDifference(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "Example 1",
			num:  11891,
			want: 99009,
		},
		{
			name: "Example 2",
			num:  90,
			want: 99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minMaxDifference(tt.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minMaxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
