package main

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
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

	expected := []int{2, 3, 4, 5, 6, 7}
	arr := inorderTraversal(tree)

	if !reflect.DeepEqual(expected, arr) {
		t.Fatalf("expected: %v, got: %v", expected, arr)
	}
}
