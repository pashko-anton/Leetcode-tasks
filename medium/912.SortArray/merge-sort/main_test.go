package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type test struct {
		nums []int
		want []int
	}

	tests := []test{
		{
			nums: []int{5, 2, 3, 1},
			want: []int{1, 2, 3, 5},
		},
		{
			nums: []int{5, 1, 1, 2, 0, 0},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tc := range tests {
		got := sortArray(tc.nums)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
