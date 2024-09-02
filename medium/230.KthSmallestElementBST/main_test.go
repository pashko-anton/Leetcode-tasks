package main

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
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

	expected := 3
	number := kthSmallest(tree, 2)

	if expected != number {
		t.Fatalf("expected: %v, got: %v", expected, number)
	}
}

func TestKthSmallestOneVal(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
	}

	expected := 1
	number := kthSmallest(tree, 1)

	if expected != number {
		t.Fatalf("expected: %v, got: %v", expected, number)
	}
}
