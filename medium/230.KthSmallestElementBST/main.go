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

func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	dfs(&arr, root)

	if len(arr) >= k {
		return arr[k-1]
	}
	return -1
}

func dfs(arr *[]int, node *TreeNode) {
	if node != nil {
		dfs(arr, node.Left)
		*arr = append(*arr, node.Val)
		dfs(arr, node.Right)
	}

	return
}
