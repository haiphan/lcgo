package problems

import (
	"reflect"
	"testing"
)

func TestMaxFreqSum(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "successes",
			want: 6,
		},
		{
			name: "Example 2",
			s:    "aeiaeia",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxFreqSum(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxFreqSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
