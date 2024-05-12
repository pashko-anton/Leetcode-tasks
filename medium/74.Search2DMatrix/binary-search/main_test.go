package main

import (
	"reflect"
	"testing"
)

func TestKthLargestElement(t *testing.T) {
	type test struct {
		nums   [][]int
		target int
		want   bool
	}

	tests := []test{
		{
			nums:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			nums:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			nums:   [][]int{{1}},
			target: 1,
			want:   true,
		},
		{
			nums:   [][]int{{1, 3}},
			target: 3,
			want:   true,
		},
	}

	for _, tc := range tests {
		got := searchMatrix(tc.nums, tc.target)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
