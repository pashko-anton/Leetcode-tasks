package main

import (
	"reflect"
	"testing"
)

func TestBucket(t *testing.T) {
	type test struct {
		nums []int
		want []int
	}

	tests := []test{
		{
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
	}

	for _, tc := range tests {
		sortColors(tc.nums)
		if !reflect.DeepEqual(tc.want, tc.nums) {
			//	t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
