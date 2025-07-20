package problems

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicateFolder(t *testing.T) {
	tests := []struct {
		name  string
		paths [][]string
		want  [][]string
	}{
		{
			name:  "Example 1",
			paths: [][]string{{"a"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}},
			want:  [][]string{{"d"}, {"d", "a"}},
		},
		{
			name:  "Example 2",
			paths: [][]string{{"a"}, {"c"}, {"a", "b"}, {"c", "b"}, {"a", "b", "x"}, {"a", "b", "x", "y"}, {"w"}, {"w", "y"}},
			want:  [][]string{{"a"}, {"a", "b"}, {"c"}, {"c", "b"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicateFolder(tt.paths)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicateFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
