package example_binary_tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	inorderTraversal(root.Left)
	fmt.Println(root.Val)
	inorderTraversal(root.Right)

	return nil
}

func preorderTraversal(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	fmt.Println(root.Val)
	inorderTraversal(root.Left)
	inorderTraversal(root.Right)

	return nil
}

func postOrderTraversal(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	inorderTraversal(root.Left)
	inorderTraversal(root.Right)
	fmt.Println(root.Val)

	return nil
}

func reverseOrderTraversal(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	reverseOrderTraversal(root.Right)
	fmt.Println(root.Val)
	reverseOrderTraversal(root.Left)

	return nil
}
