package main

import (
	"reflect"
	"testing"
)

func TestNumOfSubarrays(t *testing.T) {
	type test struct {
		nums      []int
		k         int
		threshold int
		want      int
	}

	tests := []test{
		{
			nums:      []int{2, 2, 2, 2, 5, 5, 5, 8},
			k:         3,
			threshold: 4,
			want:      3,
		},
		{
			nums:      []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3},
			k:         3,
			threshold: 5,
			want:      6,
		},
		{
			nums:      []int{99, 99},
			k:         2,
			threshold: 5,
			want:      1,
		},
	}

	for _, tc := range tests {
		got := numOfSubarrays(tc.nums, tc.k, tc.threshold)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
