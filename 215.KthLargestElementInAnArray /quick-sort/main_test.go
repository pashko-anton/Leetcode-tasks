package main

import (
	"reflect"
	"testing"
)

func TestKthLargestElement(t *testing.T) {
	type test struct {
		nums []int
		k    int
		want int
	}

	tests := []test{
		{
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
	}

	for _, tc := range tests {
		got := findKthLargest(tc.nums, tc.k)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
