package problems

import (
	"reflect"
	"testing"
)

func TestMaximumHappinessSum(t *testing.T) {
	tests := []struct {
		name      string
		happiness []int
		k         int
		want      int64
	}{
		{
			name:      "Example 1",
			happiness: []int{1, 2, 3},
			k:         2,
			want:      4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumHappinessSum(tt.happiness, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumHappinessSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
