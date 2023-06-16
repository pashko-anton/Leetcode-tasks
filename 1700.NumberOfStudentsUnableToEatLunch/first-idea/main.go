package main

import "fmt"

/*
The school cafeteria offers circular and square sandwiches at lunch break, referred to by numbers 0 and 1 respectively. All students stand in a queue. Each student either prefers square or circular sandwiches.

The number of sandwiches in the cafeteria is equal to the number of students. The sandwiches are placed in a stack. At each step:

If the student at the front of the queue prefers the sandwich on the top of the stack, they will take it and leave the queue.
Otherwise, they will leave it and go to the queue's end.
This continues until none of the queue students want to take the top sandwich and are thus unable to eat.

You are given two integer arrays students and sandwiches where sandwiches[i] is the type of the i​​​​​​th sandwich in the stack (i = 0 is the top of the stack) and students[j] is the preference of the j​​​​​​th student in the initial queue (j = 0 is the front of the queue). Return the number of students that are unable to eat.

Example 1:


Input: students = [1,1,0,0], sandwiches = [0,1,0,1]
Output: 0
Explanation:
- Front student leaves the top sandwich and returns to the end of the line making students = [1,0,0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [0,0,1,1].
- Front student takes the top sandwich and leaves the line making students = [0,1,1] and sandwiches = [1,0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [1,1,0].
- Front student takes the top sandwich and leaves the line making students = [1,0] and sandwiches = [0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [0,1].
- Front student takes the top sandwich and leaves the line making students = [1] and sandwiches = [1].
- Front student takes the top sandwich and leaves the line making students = [] and sandwiches = [].
Hence all students are able to eat.
*/

type Queue struct {
	node  *node
	last  *node
	first *node
}

type node struct {
	next *node
	prev *node
	val  int
}

func (q *Queue) BuildQueue(students []int) {
	//node = q.node
	for _, student := range students {
		if q.first == nil {
			node := &node{
				next: nil,
				prev: nil,
				val:  student,
			}
			q.first, q.last, q.node = node, node, node
		} else {
			node := &node{
				next: q.node,
				val:  student,
			}
			q.node = node
		}
	}
}

func main() {
	CountStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1})
}

func CountStudents(students []int, sandwiches []int) int {
	queue := &Queue{
		last:  nil,
		first: nil,
	}
	queue.BuildQueue(students)

	var j = 0
	for queue.first != nil {
		if queue.first.val == sandwiches[j] {
			queue.first = queue.first.next
			j++
			continue
		}

		queue.last = queue.first
		queue.first = queue.first.next
		queue.last.next = nil

	}

	fmt.Println(j)
	return 0
}
