package main

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {

	type test struct {
		students   []int
		sandwiches []int
		want       int
	}

	tests := []test{
		{
			students:   []int{1, 1, 0, 0},
			sandwiches: []int{0, 1, 0, 1},
			want:       0,
		},
		{
			students:   []int{0, 0, 0, 0},
			sandwiches: []int{1, 1, 1, 1},
			want:       4,
		},
		{
			students:   []int{1, 1, 1, 0, 0, 1},
			sandwiches: []int{1, 0, 0, 0, 1, 1},
			want:       3,
		},
	}

	for _, tc := range tests {
		got := CountStudents(tc.students, tc.sandwiches)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc)
		}
	}
}
