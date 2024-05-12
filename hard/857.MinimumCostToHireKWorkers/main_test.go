package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type test struct {
		quality []int
		wage    []int
		k       int
		want    float64
	}

	tests := []test{
		{
			quality: []int{10, 20, 5},
			wage:    []int{70, 50, 30},
			k:       2,
			want:    105,
		},
		{
			quality: []int{10},
			wage:    []int{70},
			k:       1,
			want:    70,
		},
	}

	for _, tc := range tests {
		got := mincostToHireWorkers(tc.quality, tc.wage, tc.k)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.quality)
		}
	}
}
