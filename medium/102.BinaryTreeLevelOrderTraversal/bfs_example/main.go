package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func (d *Deque) PopRight() *TreeNode {
	node := d.node[len(d.node)-1]
	d.node = d.node[:len(d.node)-1]
	return node
}

func (d *Deque) Len() int {
	return len(d.node)
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}

	deque := Deque{}
	deque.Push(root)

	level := 0
	for deque.Len() > 0 {
		length := deque.Len()
		for i := 0; i < length; i++ {
			node := deque.PopLeft()
			fmt.Println(node.Val)
			if node.Left != nil {
				deque.Push(node.Left)
			}
			if node.Right != nil {
				deque.Push(node.Right)
			}
			fmt.Println("len:", deque.Len())

		}
		level++

	}

	fmt.Println("level: ", level)
}
