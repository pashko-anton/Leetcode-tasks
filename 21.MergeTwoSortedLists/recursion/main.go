package main

/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]

*/

type LinkedList struct {
	Next *LinkedList
	Val  int
}

func main() {
}

func MergeTwoLists(input1, input2 *LinkedList) *LinkedList {
	if input1 == nil {
		return input2
	}

	if input2 == nil {
		return input1
	}

	if input1 == nil && input2 == nil {
		return nil
	}

	if input1.Val < input2.Val {
		input1.Next = MergeTwoLists(input1.Next, input2)
		return input1
	}

	input2.Next = MergeTwoLists(input1, input2.Next)

	return input2
}
