package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]
Output: [1,3,2]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]

*/

var arr []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return arr
	}

	arr := make([]int, 0)
	var inorderFillingFunc func(root *TreeNode)
	inorderFillingFunc = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorderFillingFunc(root.Left)
		arr = append(arr, root.Val)
		inorderFillingFunc(root.Right)
	}
	inorderFillingFunc(root)

	return arr
}
