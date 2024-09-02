package main

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	test1Tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	res := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := deleteNode(test1Tree, 3)
	if !reflect.DeepEqual(result, res) {
		t.Fatalf("expected: %v, got: %v", res, result)
	}

	return
}
