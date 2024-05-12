package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type input struct {
		nums []int
		n    int
	}

	type test struct {
		input input
		want  []int
	}

	tests := []test{
		{
			input: input{
				nums: []int{2, 5, 1, 3, 4, 7},
				n:    3,
			},
			want: []int{2, 3, 5, 4, 1, 7},
		},
		{
			input: input{
				nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
				n:    4,
			},
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			input: input{
				nums: []int{1, 1, 2, 2},
				n:    2,
			},
			want: []int{1, 2, 1, 2},
		},
	}

	for _, tc := range tests {
		got := Shuffle(tc.input.nums, tc.input.n)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
