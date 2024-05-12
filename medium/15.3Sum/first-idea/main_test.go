package main

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	type test struct {
		input []int
		want  [][]int
	}

	tests := []test{
		{
			input: []int{0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
	}

	for _, tc := range tests {
		got := ThreeSum(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
