package main

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {

	type test struct {
		input []int
		want  []int
	}

	tests := []test{
		{
			input: []int{1, 2, 1},
			want:  []int{1, 2, 1, 1, 2, 1},
		},
		{
			input: []int{1, 3, 2, 1},
			want:  []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		got := GetConcatenation(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
