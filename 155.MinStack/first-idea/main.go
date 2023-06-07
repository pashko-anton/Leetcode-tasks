package main

import (
	"fmt"
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
	node   *node
	length int
}

type node struct {
	prev *node
	val  int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	node := &node{
		prev: this.node,
		val:  val,
	}

	this.node = node
	this.length++

	return
}

func (this *MinStack) Pop() {
	this.node = this.node.prev
	this.length--
}

func (this *MinStack) Top() int {
	return this.node.val
}

func (this *MinStack) GetMin() int {
	var min *int
	prev := this.node
	k := 0
	for k < this.length {
		if min != nil && prev.val < *min {
			min = &prev.val
		} else if min == nil {
			min = &this.node.val
		}
		k++

		prev = prev.prev
	}

	return *min
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
