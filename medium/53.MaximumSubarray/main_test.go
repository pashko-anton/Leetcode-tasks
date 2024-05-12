package main

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	tests := []test{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			nums: []int{1},
			want: 1,
		},
		{
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			nums: []int{-3, -2, 0, -1},
			want: 0,
		},
	}

	for _, tc := range tests {
		got := maxSubArray(tc.nums)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
