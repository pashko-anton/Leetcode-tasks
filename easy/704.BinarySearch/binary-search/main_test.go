package main

import (
	"reflect"
	"testing"
)

func TestKthLargestElement(t *testing.T) {
	type test struct {
		nums   []int
		target int
		want   int
	}

	tests := []test{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -2,
			want:   -1,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 12,
			want:   5,
		},
		{
			nums:   []int{-1},
			target: -1,
			want:   0,
		},
		{
			nums:   []int{0, 1},
			target: 1,
			want:   1,
		},
	}

	for _, tc := range tests {
		got := search(tc.nums, tc.target)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
