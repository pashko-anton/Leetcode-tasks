package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	type test struct {
		input string
		want  bool
	}

	tests := []test{
		{input: "()", want: true},
		{input: "()[]{}", want: true},
		{input: "(]", want: false},
		{input: "([)]", want: false},
		{input: "[", want: false},
		{input: "[([]])", want: false},
	}

	for _, tc := range tests {
		fmt.Println(tc.input)
		got := ValidParentheses(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.input)
		}
	}
}
