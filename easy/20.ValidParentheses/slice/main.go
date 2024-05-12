package main

import (
	"fmt"
)

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

*/

func main() {
	fmt.Println(ValidParentheses("()"))
}

func ValidParentheses(s string) bool {
	stack := make([]rune, 0)
	compare := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, symbol := range s {
		if _, ok := compare[symbol]; ok {
			stack = append(stack, symbol)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if symbol == compare[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}

	return len(stack) == 0
}
