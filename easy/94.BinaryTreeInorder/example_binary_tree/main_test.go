package example_binary_tree

import (
	"fmt"
	"testing"
)

func TestBinaryTreeExampleForAnyType(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
			Left: &TreeNode{
				Val: 5,
			},
		},
	}

	fmt.Println("-----INORDER-----")
	inorderTraversal(tree)
	fmt.Println("-----------------")

	fmt.Println("-----PREORDER-----")
	preorderTraversal(tree)
	fmt.Println("-----------------")

	return
}
