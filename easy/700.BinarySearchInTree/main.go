package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

Example 1:
Input: root = [4,2,7,1,3], val = 5
Output: []
*/

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	} else if val < root.Val {
		return searchBST(root.Left, val)
	}

	return root
}
