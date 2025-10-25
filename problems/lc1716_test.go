package problems

import (
	"reflect"
	"testing"
)

func TestTotalMoney(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    4,
			want: 10,
		},
		{
			name: "Example 2",
			n:    10,
			want: 37,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalMoney(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("totalMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
