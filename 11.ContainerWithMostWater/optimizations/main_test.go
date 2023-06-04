package main

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:  49,
		},
		{
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 1},
			want:  40,
		},
		{
			input: []int{1, 1},
			want:  1,
		},
	}

	for _, tc := range tests {
		got := MaxArea(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
