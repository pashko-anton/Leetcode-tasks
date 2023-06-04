package main

import (
	"testing"
)

func TestMerge(t *testing.T) {
	type input struct {
		nums1 []int
		nums2 []int
		m     int
		n     int
	}

	type test struct {
		input input
	}

	tests := []test{
		{
			input: input{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
		{
			input: input{
				nums1: []int{4, 5, 6, 0, 0, 0},
				m:     3,
				nums2: []int{1, 2, 3},
				n:     3,
			},
		},
		{
			input: input{
				nums1: []int{-1, 0, 0, 3, 3, 3, 0, 0, 0},
				m:     6,
				nums2: []int{1, 2, 2},
				n:     3,
			},
		},
		{
			input: input{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
		},
		{
			input: input{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
	}

	for _, tc := range tests {
		Merge(tc.input.nums1, tc.input.m, tc.input.nums2, tc.input.n)
		//if !reflect.DeepEqual(tc.want, got) {
		//	t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		//}
	}
}
