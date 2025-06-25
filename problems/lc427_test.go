package problems

import (
	"reflect"
	"testing"
)

func TestConstruct(t *testing.T) {
	t.Run("test quad tree", func(t *testing.T) {
		grid := [][]int{
			{0, 1},
			{1, 0},
		}
		got := construct(grid)
		if !reflect.DeepEqual(got.IsLeaf, false) {
			t.Errorf("construct() = %v, want %v", got.IsLeaf, false)
		}
		if !reflect.DeepEqual(got.TopRight.IsLeaf, true) {
			t.Errorf("construct() = %v, want %v", got.TopRight.IsLeaf, true)
		}
	})
}
