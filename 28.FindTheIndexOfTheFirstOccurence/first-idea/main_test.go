package main

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	type test struct {
		haystack string
		needle   string
		want     int
	}

	tests := []test{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "sadbutsad",
			needle:   "but",
			want:     3,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
		{
			haystack: "mississippi",
			needle:   "issipi",
			want:     -1,
		},
		{
			haystack: "abc",
			needle:   "c",
			want:     2,
		},
	}

	for _, tc := range tests {
		got := StrStr(tc.haystack, tc.needle)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v str: %v", tc.want, got, tc.haystack, tc.needle)
		}
	}
}
