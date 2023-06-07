package main

import (
	"fmt"
	"math"
)

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.


Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
*/

type MinStack struct {
	head *node
}

type node struct {
	prev *node
	val  int
	min  int
}

func Constructor() MinStack {
	return MinStack{
		head: &node{min: math.MaxInt32},
	}
}

func (this *MinStack) Push(val int) {
	min := val
	if val > this.head.min {
		min = this.head.min
	}
	this.head = &node{
		prev: this.head,
		val:  val,
		min:  min,
	}
	return
}

func (this *MinStack) Pop() {
	this.head = this.head.prev
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Top()
	fmt.Println(minStack.GetMin())
}
