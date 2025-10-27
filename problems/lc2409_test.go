package problems

import (
	"reflect"
	"testing"
)

func TestCountDaysTogether(t *testing.T) {
	tests := []struct {
		name                                         string
		arriveAlice, leaveAlice, arriveBob, leaveBob string
		want                                         int
	}{
		{
			name:        "Example 1",
			arriveAlice: "08-15",
			leaveAlice:  "08-18",
			arriveBob:   "08-16",
			leaveBob:    "08-19",
			want:        3,
		},
		{
			name:        "Example 2",
			arriveAlice: "10-01",
			leaveAlice:  "10-31",
			arriveBob:   "11-01",
			leaveBob:    "12-31",
			want:        0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countDaysTogether(tt.arriveAlice, tt.leaveAlice, tt.arriveBob, tt.leaveBob)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countDaysTogether() = %v, want %v", got, tt.want)
			}
		})
	}
}
