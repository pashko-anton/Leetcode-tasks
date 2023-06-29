package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	tests := []test{
		{
			nums: []int{1, -2, 3, -2},
			want: 3,
		},
		{
			nums: []int{5, -3, 5},
			want: 10,
		},
		{
			nums: []int{-3, -2, -3},
			want: -2,
		},
	}

	for _, tc := range tests {
		got := maxSubarraySumCircular(tc.nums)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
