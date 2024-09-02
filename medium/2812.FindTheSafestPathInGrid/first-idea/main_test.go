package first_idea

import (
	"reflect"
	"testing"
)

func TestMaximumSafenessFactor(t *testing.T) {
	type test struct {
		grid [][]int
		want int
	}

	tests := []test{
		{
			grid: [][]int{
				{0, 0, 1},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 2,
		},
		{
			grid: [][]int{
				{0, 0, 0, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 0, 0, 0},
			},
			want: 2,
		},
		{
			grid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			want: 0,
		},
	}

	for _, tc := range tests {
		got := maximumSafenessFactor(tc.grid)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.grid)
		}
	}

	return
}
