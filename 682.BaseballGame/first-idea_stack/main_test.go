package main

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {

	type test struct {
		input []string
		want  int
	}

	tests := []test{
		{
			input: []string{"5", "2", "C", "D", "+"},
			want:  30,
		},
		{
			input: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			want:  27,
		},
		{
			input: []string{"1", "C"},
			want:  0,
		},
	}

	for _, tc := range tests {
		got := CalPoints(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
