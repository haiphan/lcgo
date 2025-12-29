package problems

import (
	"reflect"
	"testing"
)

func TestCountSeniors(t *testing.T) {
	tests := []struct {
		name    string
		details []string
		want    int
	}{
		{
			name:    "Example 1",
			details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"},
			want:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSeniors(tt.details)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSeniors() = %v, want %v", got, tt.want)
			}
		})
	}
}
