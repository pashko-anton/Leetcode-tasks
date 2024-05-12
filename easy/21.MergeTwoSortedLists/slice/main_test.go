package main

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	type slices struct {
		list1 []int
		list2 []int
	}
	type test struct {
		input slices
		want  []int
	}

	tests := []test{
		{
			input: slices{
				list1: []int{1, 2},
				list2: []int{1, 2},
			},
			want: []int{1, 1, 2, 2},
		},
		{
			input: slices{
				list1: []int{1, 2, 4},
				list2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input: slices{
				list1: []int{},
				list2: []int{},
			},
			want: []int{},
		},
		{
			input: slices{
				list1: []int{},
				list2: []int{0},
			},
			want: []int{0},
		},
	}

	for _, tc := range tests {
		got := MergeTwoLists(tc.input.list1, tc.input.list2)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
