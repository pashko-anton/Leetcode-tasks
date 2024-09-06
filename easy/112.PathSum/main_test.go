package main

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
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

	isHas := hasPathSum(tree, 9)

	if !reflect.DeepEqual(isHas, true) {
		t.Fatalf("expected: %v, got: %v", true, isHas)
	}
}

func TestPathSum2(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	isHas := hasPathSum(tree, 22)

	if !reflect.DeepEqual(isHas, true) {
		t.Fatalf("expected: %v, got: %v", true, isHas)
	}
}
