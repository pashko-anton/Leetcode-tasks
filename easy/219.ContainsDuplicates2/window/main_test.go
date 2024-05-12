package main

import (
	"reflect"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	type test struct {
		nums []int
		k    int
		want bool
	}

	tests := []test{
		{
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
		{
			nums: []int{99, 99},
			k:    2,
			want: true,
		},
		{
			nums: []int{2, 2},
			k:    3,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9},
			k:    3,
			want: true,
		},
	}

	for _, tc := range tests {
		got := containsNearbyDuplicate(tc.nums, tc.k)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
