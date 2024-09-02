package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.


Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
Example 2:

Input: root = [1,null,3]
Output: [1,3]
Example 3:

Input: root = []
Output: []

*/

type Deque struct {
	node []*TreeNode
}

func (d *Deque) Push(node *TreeNode) {
	d.node = append(d.node, node)
}

func (d *Deque) PopLeft() *TreeNode {
	node := d.node[0]
	d.node = d.node[1:]
	return node
}

func (d *Deque) Len() int {
	return len(d.node)
}

func rightSideView(root *TreeNode) []int {
	var resp = make([]int, 0)

	if root == nil {
		return []int{}
	}

	deque := Deque{}
	deque.Push(root)

	for deque.Len() > 0 {
		lastVal := 0
		length := deque.Len()
		for i := 0; i < length; i++ {
			node := deque.PopLeft()
			lastVal = node.Val
			if node.Left != nil {
				deque.Push(node.Left)
			}
			if node.Right != nil {
				deque.Push(node.Right)
			}
		}
		resp = append(resp, lastVal)
	}

	return resp
}
