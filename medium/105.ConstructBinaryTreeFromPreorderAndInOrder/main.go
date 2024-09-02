package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]

*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootNode := &TreeNode{
		Val: preorder[0],
	}
	idx := slices.Index(inorder, preorder[0])
	rootNode.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	rootNode.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return rootNode
}
