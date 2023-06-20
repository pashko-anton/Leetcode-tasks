package main

import (
	"reflect"
	"testing"
)

func TestKthLargestElement(t *testing.T) {
	type test struct {
		nums  []int
		hours int
		want  int
	}

	tests := []test{
		{
			nums:  []int{3, 6, 7, 11},
			hours: 8,
			want:  4,
		},
		{
			nums:  []int{30, 11, 23, 4, 20},
			hours: 5,
			want:  30,
		},
		{
			nums:  []int{30, 11, 23, 4, 20},
			hours: 6,
			want:  23,
		},
		{
			nums:  []int{312884470},
			hours: 968709470,
			want:  1,
		},
	}

	for _, tc := range tests {
		got := minEatingSpeed(tc.nums, tc.hours)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
