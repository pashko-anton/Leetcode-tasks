package main

import "fmt"

/*
Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node.
If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement the MyLinkedList class:

MyLinkedList() Initializes the MyLinkedList object.
int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
void addAtTail(int val) Append a node of value val as the last element of the linked list.
void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.



Example 1:

Input
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
Output
[null, null, null, null, 2, null, 3]

Explanation
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
myLinkedList.get(1);              // return 2
myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
myLinkedList.get(1);              // return 3
*/

type MyLinkedList struct {
	node *node
	head *node
	tail *node
}

type node struct {
	prev *node
	next *node
	val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		index = 0
	}
	var i int
	node := this.node
	for node != nil {
		if index == i {
			return node.val
		}
		node = node.next
		i++
	}

	return -1
}

func (this *MyLinkedList) Counter() int {

	var i int
	node := this.node
	for node != nil {
		node = node.next
		i++
	}

	return i
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.node == nil {
		this.node = &node{
			prev: nil,
			next: nil,
			val:  val,
		}
		this.head = this.node
		this.tail = this.node
		return
	}

	prev := this.tail
	tailNode := &node{
		prev: prev,
		next: nil,
		val:  val,
	}
	this.tail, prev.next = tailNode, tailNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	}

	if this.node == nil && index == 0 {
		this.node = &node{
			prev: nil,
			next: nil,
			val:  val,
		}
		this.head = this.node
		this.tail = this.node
		return
	}

	var i int
	nodeCursor := this.node
	for nodeCursor != nil {
		if index == i {
			newNode := &node{
				prev: nodeCursor.prev,
				next: nodeCursor,
				val:  val,
			}

			if i == 0 {
				this.head = newNode
				this.node = newNode
				nodeCursor.prev = newNode
			} else {
				nodeCursor.prev.next = newNode
				nodeCursor.prev = newNode
			}
			return
		}

		i++

		if nodeCursor.next == nil && index == i {
			nodeCursor.next = &node{
				prev: nodeCursor,
				next: nil,
				val:  val,
			}
			this.tail = nodeCursor.next
			return
		}
		nodeCursor = nodeCursor.next

	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		index = 0
	}

	if index == 0 {
		this.node = this.node.next
		this.head = this.node
	}
	var i int

	node := this.node
	for node != nil {
		if index == i && index != 0 {
			node.prev.next = node.next
			if node.next == nil {
				this.tail = node.prev
			} else {
				node.next.prev = node.prev
			}
			return
		}

		node = node.next
		i++
	}
}

// 1, 2, 7
func main() {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(7)
	myLinkedList.AddAtHead(2)
	myLinkedList.AddAtHead(1)

	myLinkedList.AddAtIndex(3, 0)
	myLinkedList.DeleteAtIndex(2)

	myLinkedList.DeleteAtIndex(2)
	myLinkedList.AddAtHead(6)
	myLinkedList.AddAtTail(4)
	fmt.Println(myLinkedList.Get(4))
	myLinkedList.AddAtHead(4)
	myLinkedList.AddAtIndex(5, 0)
	myLinkedList.AddAtHead(6)
}
