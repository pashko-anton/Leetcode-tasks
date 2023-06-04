package main

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{
			input: []int{1},
			want:  1,
		},
		{
			input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:  5,
		},
	}

	for _, tc := range tests {
		got := RemoveDuplicates(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
