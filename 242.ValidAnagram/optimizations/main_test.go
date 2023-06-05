package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type input struct {
		s string
		t string
	}

	type test struct {
		input input
		want  bool
	}

	tests := []test{
		{
			input: input{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			input: input{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			input: input{
				s: "aaaaa",
				t: "aaaaaa",
			},
			want: false,
		},
		{
			input: input{
				s: "x",
				t: "x",
			},
			want: true,
		},
	}

	for _, tc := range tests {
		got := IsAnagram(tc.input.s, tc.input.t)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
