package problems

import (
	"reflect"
	"testing"
)

func TestNumberOfStableArrays(t *testing.T) {
	tests := []struct {
		name                   string
		zero, one, limit, want int
	}{
		{
			name:  "Example 1",
			zero:  1,
			one:   1,
			limit: 2,
			want:  2,
		},
		{
			name:  "Example 2",
			zero:  1,
			one:   2,
			limit: 1,
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfStableArrays(tt.zero, tt.one, tt.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfStableArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
