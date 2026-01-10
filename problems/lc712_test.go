package problems

import (
	"reflect"
	"testing"
)

func TestMinimumDeleteSum(t *testing.T) {
	tests := []struct {
		name   string
		s1, s2 string
		want   int
	}{
		{
			name: "Example 1",
			s1:   "sea",
			s2:   "eat",
			want: 231,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumDeleteSum(tt.s1, tt.s2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
