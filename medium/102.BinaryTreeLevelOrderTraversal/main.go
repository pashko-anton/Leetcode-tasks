package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

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

func levelOrder(root *TreeNode) [][]int {
	var resp = make([][]int, 0)

	if root == nil {
		return resp
	}

	deque := Deque{}
	deque.Push(root)

	level := 0
	for deque.Len() > 0 {
		resp = append(resp, []int{})

		length := deque.Len()
		for i := 0; i < length; i++ {
			node := deque.PopLeft()
			resp[level] = append(resp[level], node.Val)

			if node.Left != nil {
				deque.Push(node.Left)
			}
			if node.Right != nil {
				deque.Push(node.Right)
			}
		}
		level++
	}

	return resp
}
