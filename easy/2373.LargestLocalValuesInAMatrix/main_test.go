package main

import (
	"reflect"
	"testing"
)

func TestLargestLocal(t *testing.T) {
	type test struct {
		grid [][]int
		want [][]int
	}

	tests := []test{
		{
			grid: [][]int{
				{9, 9, 8, 1},
				{5, 6, 2, 6},
				{8, 2, 6, 4},
				{6, 2, 2, 2},
			},
			want: [][]int{
				{9, 9},
				{8, 6},
			},
		},
		{
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 2, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
		},
	}

	for _, tc := range tests {
		got := largestLocal(tc.grid)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.grid)
		}
	}

	return
}
