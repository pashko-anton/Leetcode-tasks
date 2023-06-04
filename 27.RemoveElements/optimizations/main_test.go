package main

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	type test struct {
		input     []int
		duplicate int
		want      int
	}

	tests := []test{
		{
			input:     []int{3, 2, 2, 3},
			duplicate: 2,
			want:      2,
		},
		{
			input:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			duplicate: 2,
			want:      5,
		},
	}

	for _, tc := range tests {
		got := RemoveDuplicates(tc.input, tc.duplicate)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
