package main

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	expected := []int{1, 3, 4}
	resp := rightSideView(tree)

	if !reflect.DeepEqual(expected, resp) {
		t.Fatalf("expected: %v, got: %v", expected, resp)
	}
}
