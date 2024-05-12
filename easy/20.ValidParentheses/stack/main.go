package main

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

type Stack struct {
	top    *node
	length int
}

type node struct {
	prev  *node
	value rune
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) Peek() rune {
	if s.length == 0 {
		return 0
	}

	return s.top.value
}

func (s *Stack) Push(value rune) {
	if s.top == nil {
		s.top = &node{
			prev:  nil,
			value: value,
		}
	}

	newNode := &node{prev: s.top, value: value}
	s.top = newNode
	s.length += 1
}

func (s *Stack) Pop() rune {
	if s.length == 0 {
		return 0
	}

	val := s.top.value
	s.top = s.top.prev
	s.length -= 1
	return val
}

func main() {
	ValidParentheses("()")
}

func ValidParentheses(s string) bool {
	var stack Stack
	compare := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, symbol := range s {
		if _, ok := compare[symbol]; ok {
			stack.Push(symbol)
			continue
		}

		if stack.Length() == 0 || compare[stack.Pop()] != symbol {
			return false
		}
	}

	return stack.Length() == 0
}
