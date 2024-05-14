package first_idea

import (
	"reflect"
	"testing"
)

func TestGetMaximumGold(t *testing.T) {
	type test struct {
		grid [][]int
		want int
	}

	tests := []test{
		{
			grid: [][]int{
				{0, 6, 0},
				{5, 8, 7},
				{0, 9, 0},
			},
			want: 24,
		},
		{
			grid: [][]int{
				{1, 0, 7},
				{2, 0, 6},
				{3, 4, 5},
				{0, 3, 0},
				{9, 0, 20},
			},
			want: 28,
		},
	}

	for _, tc := range tests {
		got := getMaximumGold(tc.grid)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.grid)
		}
	}

	return
}
