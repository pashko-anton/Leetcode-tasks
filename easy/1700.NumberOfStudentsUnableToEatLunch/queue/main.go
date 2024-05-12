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
	node *node
	last *node
}

type node struct {
	prev *node
	val  int
}

func (this *Queue) Empty() bool {
	return this.node == nil
}

func (this *Queue) BuildQueue(students []int) {
	for i := len(students) - 1; i >= 0; i-- {
		if this.Empty() {
			this.node = &node{
				prev: nil,
				val:  students[i],
			}
			this.last = this.node
			continue
		}

		prev := this.node
		this.node = &node{
			prev: prev,
			val:  students[i],
		}
	}
}

func main() {
	fmt.Println(CountStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
}

func CountStudents(students []int, sandwiches []int) int {
	queue := &Queue{}
	queue.BuildQueue(students)

	var j = 0

	for queue.node.prev != nil {
		if queue.node.val == sandwiches[j] {
			queue.node = queue.node.prev
			j++
			continue
		}

		queue.last.prev = queue.node
		queue.node = queue.node.prev
		queue.last = queue.last.prev
		queue.last.prev = nil
	}

	return len(sandwiches) - (j + 1)
}
