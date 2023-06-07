package main

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ReverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})
}

func ReverseList(head *ListNode) *ListNode {
	var listNode *ListNode
	for head != nil {
		head.Next, head, listNode = listNode, head.Next, head
	}

	return listNode
}
