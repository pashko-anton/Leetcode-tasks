package main

import "testing"

func TestBFS(t *testing.T) {
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

	bfs(tree)

}
