package main

import "strconv"

/*
You are keeping the scores for a baseball game with strange rules. At the beginning of the game, you start with an empty record.

You are given a list of strings operations, where operations[i] is the ith operation you must apply to the record and is one of the following:

An integer x.
Record a new score of x.
'+'.
Record a new score that is the sum of the previous two scores.
'D'.
Record a new score that is the double of the previous score.
'C'.
Invalidate the previous score, removing it from the record.
Return the sum of all the scores on the record after applying all the operations.

The test cases are generated such that the answer and all intermediate calculations fit in a 32-bit integer and that all operations are valid.

Example 1:

Input: ops = ["5","2","C","D","+"]
Output: 30
Explanation:
"5" - Add 5 to the record, record is now [5].
"2" - Add 2 to the record, record is now [5, 2].
"C" - Invalidate and remove the previous score, record is now [5].
"D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
"+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
The total sum is 5 + 10 + 15 = 30.

Example 2:
Input: ops = ["5","-2","4","C","D","9","+","+"]
Output: 27
Explanation:
"5" - Add 5 to the record, record is now [5].
"-2" - Add -2 to the record, record is now [5, -2].
"4" - Add 4 to the record, record is now [5, -2, 4].
"C" - Invalidate and remove the previous score, record is now [5, -2].
"D" - Add 2 * -2 = -4 to the record, record is now [5, -2, -4].
"9" - Add 9 to the record, record is now [5, -2, -4, 9].
"+" - Add -4 + 9 = 5 to the record, record is now [5, -2, -4, 9, 5].
"+" - Add 9 + 5 = 14 to the record, record is now [5, -2, -4, 9, 5, 14].
The total sum is 5 + -2 + -4 + 9 + 5 + 14 = 27.

Example 3:
Input: ops = ["1","C"]
Output: 0
Explanation:
"1" - Add 1 to the record, record is now [1].
"C" - Invalidate and remove the previous score, record is now [].
Since the record is empty, the total sum is 0.
*/

type Stack struct {
	node   *node
	length int
}

type node struct {
	prev *node
	val  int
}

func (s *Stack) Peek() int {
	return s.node.val
}

func (s *Stack) Pop() int {
	val := s.node.val
	s.node = s.node.prev
	s.length--

	return val
}

func (s *Stack) Push(val int) {
	node := &node{
		prev: s.node,
		val:  val,
	}

	s.node = node
	s.length++

	return
}

func main() {
}

func CalPoints(nums []string) int {
	var stack = Stack{}
	for _, n := range nums {
		switch n {
		case "+":
			stackPopVal := stack.Pop()
			sum := stackPopVal + stack.Peek()
			stack.Push(stackPopVal)
			stack.Push(sum)
		case "D":
			stack.Push(stack.Peek() * 2)
		case "C":
			stack.Pop()
		default:
			v, _ := strconv.Atoi(n)
			stack.Push(v)
		}
	}

	var sum int
	for {
		if stack.length == 0 {
			return sum
		}
		sum += stack.Pop()
	}
}
