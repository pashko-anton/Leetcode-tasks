package main

import (
	"reflect"
	"testing"
)

func TestLevelOrderCycle(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Right: &TreeNode{
				Val: 7,
			},
			Left: &TreeNode{
				Val: 15,
			},
		},
	}

	expected := [][]int{
		0: {3},
		1: {9, 20},
		2: {15, 7},
	}
	resp := levelOrder(tree)

	if !reflect.DeepEqual(expected, resp) {
		t.Fatalf("expected: %v, got: %v", expected, resp)
	}
}
