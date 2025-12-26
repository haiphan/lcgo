package problems

import (
	"reflect"
	"testing"
)

func TestBestClosingTime(t *testing.T) {
	tests := []struct {
		name      string
		customers string
		want      int
	}{
		{
			name:      "Example 1",
			customers: "YYNY",
			want:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bestClosingTime(tt.customers)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestClosingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
