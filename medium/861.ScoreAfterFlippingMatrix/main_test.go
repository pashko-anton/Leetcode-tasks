package main

import (
	"reflect"
	"testing"
)

func TestMatrixScore(t *testing.T) {
	type test struct {
		grid [][]int
		want int
	}

	tests := []test{
		{
			grid: [][]int{
				{0, 0, 1, 1},
				{1, 0, 1, 0},
				{1, 1, 0, 0},
			},
			want: 39,
		},
		{
			grid: [][]int{
				{0},
			},
			want: 1,
		},
	}

	for _, tc := range tests {
		got := matrixScore(tc.grid)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.grid)
		}
	}

	return
}
